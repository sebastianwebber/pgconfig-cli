// Copyright © 2017 Sebastian Webber <sebastian@swebber.me>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/go-resty/resty"
	"github.com/spf13/cobra"
)

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "  ")
	if err != nil {
		return in
	}
	return out.String()
}

func CallAPI(
	pgVersion string,
	totalRAM int,
	maxConnections int,
	environmentName string,
	osType string,
	arch string,
	format string,
	includePGBadger bool,
	logFormat string,
	pretty bool) string {
	// https: //api.pgconfig.org/v1/tuning/get-config?enviroment_name=Desktop&format=json&include_pgbadger=true&log_format=stderr&max_connections=100&pg_version=9.5&total_ram=29GB

	restParms := map[string]string{
		"pg_version":       pgVersion,
		"total_ram":        strconv.Itoa(totalRAM) + "GB",
		"max_connections":  strconv.Itoa(maxConnections),
		"environment_name": environmentName,
		"os_type":          osType,
		"arch":             arch,
		"format":           format,
		"include_pgbadger": strconv.FormatBool(includePGBadger),
		"log_format":       logFormat,
	}

	// fmt.Println(restParms, pretty)

	resp, err := resty.R().
		SetQueryParams(restParms).
		Get("https://api.pgconfig.org/v1/tuning/get-config")

	if err != nil {
		log.Fatal(err)
	}

	output := resp.String()

	if pretty == true {
		output = jsonPrettyPrint(output)
	}

	return output
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the configuration from the pgconfig api",
	Long:  `Get the configuration from the pgconfig api`,
	Run: func(cmd *cobra.Command, args []string) {

		output := CallAPI(
			pgVersion,
			totalRAM,
			maxConnections,
			environmentName,
			osType,
			arch,
			format,
			includePGBadger,
			logFormat,
			pretty)

		fmt.Println(output)
	},
}

var (
	pgVersion       = "9.6"
	totalRAM        = 2
	maxConnections  = 100
	environmentName = "WEB"
	osType          = "Linux"
	arch            = "x86-64"
	format          = "json"
	includePGBadger = false
	logFormat       = "stderr"
	pretty          = false
)

func init() {

	getCmd.Flags().StringVarP(&pgVersion, "pg-version", "v", pgVersion, "PostgreSQL Version")
	getCmd.Flags().IntVarP(&totalRAM, "memory", "m", totalRAM, "Total memory to use")
	getCmd.Flags().IntVarP(&maxConnections, "max-connections", "c", maxConnections, "Max connections")
	getCmd.Flags().StringVarP(&environmentName, "environment-name", "e", environmentName, "Environment name")
	getCmd.Flags().StringVarP(&osType, "os-type", "o", osType, "Operating system type")
	getCmd.Flags().StringVarP(&arch, "arch", "a", arch, "Operating system arch")
	getCmd.Flags().StringVarP(&format, "format", "f", format, "output format")
	getCmd.Flags().BoolVarP(&includePGBadger, "include-pgbadger", "b", includePGBadger, "Include PGBadger configuration")
	getCmd.Flags().StringVarP(&logFormat, "log-format", "l", logFormat, "Log format (if included pgbadger stuff)")

	getCmd.Flags().BoolVarP(&pretty, "pretty", "P", pretty, "format output (json only)")

	RootCmd.AddCommand(getCmd)
}
