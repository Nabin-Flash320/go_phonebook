
package Settings


import (
	"net"
)


var allowedHosts = [...]string{
	"192.168.254.156",
}


func HostSettingsGetAllowedHosts() []string {

	return allowedHosts[:]

}


func HostSettingsGetValidHostAddress() string {

	addrs, error := net.InterfaceAddrs()

	if error != nil {
		
		panic(error)

	}

	for _, addr := range addrs {

        ipNet, ok := addr.(*net.IPNet)
        if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {

           if hostSettingsContainsHost(ipNet.IP.String()) {

			   return ipNet.IP.String()

		   }

        }

    }
	
	return "127.0.0.1"

}


func hostSettingsContainsHost(s string) bool {

	for _, str := range allowedHosts {

		if str == s {

			return true

		}

	}
	return false

}



