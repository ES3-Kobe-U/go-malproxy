package handler

import "net"

// IPアドレスから逆引きでホスト名を検索する関数
func GetHostnameFromIPAddress(ipaddress string) ([]string, error) {
	hostname, err := net.LookupAddr(ipaddress)
	if err != nil {
		return nil, err
	}
	return hostname, nil
}