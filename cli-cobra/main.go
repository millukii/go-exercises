//example go run main.go echo times Hello World -t 5
//esto debería imprimir por consola Echo: Hello World  5 veces
//go run main.go echo time 
//imprime una vez echo time
package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	persistRootFlag bool
	localRootFlag   bool
	times int
	rootCmd         = &cobra.Command{
		Use:   "example",
		Short: "an example cli go app with cobra",
		Long:  "This is a simple example of cli golang program",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello from root command")
		},
	}
	echoCmd = &cobra.Command{
		Use:   "echo [strings to echo]",
		Short: "prints given strings to stdout",
		//el mínimo de argumentos que requiere
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	timesCmd = &cobra.Command{
		Use:   "times [strings to echo]",
		Short: "prints given strings to stdout multiple time",
		Args: cobra.MinimumNArgs(1),
		Run: func (cmd *cobra.Command, args []string)  error{
			if times == 0{
				return errors.New("times cannot be 0")
			}
			for i := 0; i < times; i++ {
				fmt.Println("Echo: "+strings.Join(args, " "))
			}
			return nil
		}
	}
)

//cobra tiene 2 tipos de flag, las persistentes y las locales
//las locales solo estan disponibles en el contexto de definicion del comando
//las persistentes estan disponibles dentro de la definición del comando y sus subcomandos
func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persisten root flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "1".false, "a local root var")
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	timesCmd.MarkFlagRequired("times")
	rootCmd.AddCommand(echoCmd)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("go cli with cobra")
}
