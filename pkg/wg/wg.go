package wg

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ezh/wireguard-grpc/pkg/exec"
	"github.com/ezh/wireguard-grpc/pkg/utilities"
	"github.com/go-logr/logr"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var (
	ErrValueParse               = fmt.Errorf("value parse failed")
	ErrPersistentKeepaliveRange = fmt.Errorf("persistent keepalive interval is neither 0/off nor 1-65535")
)

type Exec struct {
	exec.Executor
}

func New(rawCmd string) *Exec {
	return &Exec{Executor: exec.New(rawCmd)}
}

func (exe *Exec) Verify(l *logr.Logger) bool {
	stdout, stderr, err := exe.Run(l, "show")
	if err != nil || len(stderr) > 0 {
		l.Error(err, "wg failed", "stdout", stdout, "stderr", stderr)
		return false
	}
	return true
}

func (exe *Exec) Version(l *logr.Logger) (string, error) {
	stdout, _, err := exe.Run(l, "-v")
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`\bv\d+[[:graph:]]+`)
	version := re.FindString(stdout)
	if len(version) == 0 {
		return "", errors.New("unable to get wg version")
	}
	return version, nil
}

func (exe *Exec) Dump(l *logr.Logger) ([]*wgtypes.Device, error) {
	var devices []*wgtypes.Device
	var device *wgtypes.Device

	stdout, _, err := exe.Run(l, "show", "all", "dump")
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(stdout, "\n") {
		parts := strings.Split(line, "\t")
		if len(parts) == 5 {
			// We never touch private data
			// privateKey, err := wgtypes.ParseKey(parts[1])
			// if err != nil {
			// 	return nil, err
			// }
			publicKey, err := wgtypes.ParseKey(parts[2])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: PublicKey=%s", ErrValueParse, parts[2])
			}
			port, err := parseInt64(parts[3])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: ListenPort=%s", ErrValueParse, parts[3])
			}
			fwMark, err := parseFwMark(parts[4])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: FwMark=%s", ErrValueParse, parts[4])
			}
			device = &wgtypes.Device{
				Name:         parseStr(parts[0]),
				Type:         wgtypes.Unknown,
				PrivateKey:   wgtypes.Key{},
				PublicKey:    publicKey,
				ListenPort:   int(port),
				FirewallMark: fwMark,
			}
			devices = append(devices, device)
		}
		if len(parts) == 9 && device != nil && device.Name == parts[0] {
			publicKey, err := wgtypes.ParseKey(parts[1])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: PublicKey=%s", ErrValueParse, parts[1])
			}
			presharedKey, err := wgtypes.ParseKey(parts[2])
			if err != nil && !strings.Contains(parts[2], "none") {
				return nil, errors.Wrapf(err, "%v: PresharedKey=%s", ErrValueParse, parts[2])
			}
			endpointAddr, err := net.ResolveUDPAddr("udp", parts[3])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: EndpointAddr=%s", ErrValueParse, parts[3])
			}
			allowedIPs, err := utilities.StringToIPNetSlice(parts[4])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: AllowedIPs=%v", ErrValueParse, parts[4])
			}
			lastHandshakeTime, err := parseInt64(parts[5])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: LastHandshakeTime=%v", ErrValueParse, parts[5])
			}
			receiveBytes, err := parseInt64(parts[6])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: ReceiveBytes=%v", ErrValueParse, parts[6])
			}
			transmitBytes, err := parseInt64(parts[7])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: TransmitBytes=%v", ErrValueParse, parts[7])
			}
			persistentKeepalive, err := parsePersistentKeepalive(parts[8])
			if err != nil {
				return nil, errors.Wrapf(err, "%v: PersistentKeepalive=%v", ErrValueParse, parts[8])
			}
			peer := wgtypes.Peer{
				PublicKey:                   publicKey,
				PresharedKey:                presharedKey,
				Endpoint:                    endpointAddr,
				PersistentKeepaliveInterval: persistentKeepalive,
				LastHandshakeTime:           time.Unix(lastHandshakeTime, 0),
				ReceiveBytes:                receiveBytes,
				TransmitBytes:               transmitBytes,
				AllowedIPs:                  allowedIPs,
			}
			device.Peers = append(device.Peers, peer)
		}
	}

	return devices, err
}

func parseStr(s string) string {
	if s == "" {
		return s
	}
	if s == "(none)" {
		return ""
	}
	return strings.TrimSpace(s)
}

func parseInt64(s string) (int64, error) {
	ps := parseStr(s)
	if ps == "" {
		return 0, nil
	}
	return strconv.ParseInt(ps, 10, 64)
}

func parseFwMark(s string) (int, error) {
	if strings.EqualFold(s, "off") {
		return 0, nil
	}
	fwMark, err := strconv.ParseInt(s, 0, 0)
	return int(fwMark), err
}

func parsePersistentKeepalive(s string) (time.Duration, error) {
	if strings.EqualFold(s, "off") {
		return time.Duration(0), nil
	}
	persistentKeepalive, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return time.Duration(0), err
	}
	if persistentKeepalive < 0 || persistentKeepalive > 65535 {
		return time.Duration(0), ErrPersistentKeepaliveRange
	}

	return time.Duration(persistentKeepalive * int64(time.Second)), err
}
