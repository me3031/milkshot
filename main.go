package main 

import ( 
	"fmt"

	"github.com/p0w3r-surg3/core/cli"
	"github.com/p0w3r-surg3/core/socket/server"

)

func main() {

	var context cli.Context = cli.InitFlags() 

	if context.Target.HostCanBeResolv() {
		fmt.Printf("Target resolved to %s ($s)\n", context.Target.Ipv4.String(), context.Target.Ipv6.String())

	}

	if context.UseListenMode {
		fmt.Printf("Listening on %s:%d\n", context.Target.Host, context.Target.Port)

		var server server.Server = server.Create(context.Target, context.UseDebugMode)

		server.Start()
	} else {

		fmt.Printf("Initialize Client Mode : WIP")
		
	}
}