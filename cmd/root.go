package cmd

import (
	"net/url"
	"os"

	"github.com/hackermanpeter/caching-proxy/internal/cache"
	"github.com/hackermanpeter/caching-proxy/internal/server"
	"github.com/spf13/cobra"
)

var port int
var origin string
var clearCache bool

var rootCmd = &cobra.Command{
	Use:   "caching-proxy",
	Short: "Caching Proxy it will forward requests to the actual server and cache the responses",
	Long: `Caching Proxy it will forward requests to the actual server and cache the responses. If the same 
request is made again, it will return the cached response instead of forwarding the request to the server.
For Example:

caching-proxy --port 3000 --origin http://dummyjson.com
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if clearCache {
			cache.Empty()
			return nil
		}

		_, err := url.ParseRequestURI(origin)
		if err != nil {
			return err
		}

		server.Start(port, origin)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVar(&port, "port", 3000, "Port to listen on")
	rootCmd.Flags().StringVar(&origin, "origin", "", "Url to forward requests to")
	rootCmd.Flags().BoolVar(&clearCache, "clear-cache", false, "Clear application cache")

	rootCmd.MarkFlagsOneRequired("origin", "clear-cache")
	rootCmd.MarkFlagsMutuallyExclusive("origin", "clear-cache")
}
