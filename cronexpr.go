package main

import (
  "os"
  "github.com/codegangsta/cli"
  "github.com/gorhill/cronexpr"
  "time"
  "fmt"
  "strconv"
  "github.com/hhkbp2/go-strftime"
  //"github.com/jehiah/go-strftime"
  //"github.com/tebeka/strftime"
)

func main() {
  app := cli.NewApp()
  app.Version = "0.1.0"
  app.Name = "cronexpr"
  app.Usage = "convert cron expression and get next occurance"
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "unix, u",
      Value: "",
      Usage: "from specific unix timestamp",
    },
    cli.StringFlag{
      Name: "format, f",
      Value: "",
      Usage: "format options see http://strftime.org/",
    },
    cli.StringFlag{
      Name: "next, n",
      Value: "",
      Usage: "n next time stamps",
    },
    cli.StringFlag{
      Name: "utc",
      Value: "false",
      Usage: "n next time stamps",
    },
  }
  app.Action = func(c *cli.Context) {
    cron := ""
    if len(c.Args()) > 0 {
      cron = c.Args()[0]
    } else {
      panic("missing cron expression")
    }

    from := time.Now()
    if c.String("unix") != "" {
      u, err := strconv.ParseInt(c.String("unix"), 10, 64)
      if err != nil {
          panic(err)
      }
      from = time.Unix(u, 0)
    }
    if c.BoolT("utc") {
      from = from.UTC()
    }

    if c.String("next") != "" {
      n, err := strconv.ParseInt(c.String("next"), 10, 64)
      if err != nil {
          panic(err)
      }
      result := cronexpr.MustParse(cron).NextN(from, uint(n))
      for _, next := range result {
        out := strconv.FormatInt(next.Unix(), 10)
        if c.String("format") != "" {
          out = strftime.Format(c.String("format"), next)
        }
        fmt.Println(out)
      }
    } else {
      result := cronexpr.MustParse(cron).Next(from)
      out := strconv.FormatInt(result.Unix(), 10)
      if c.String("format") != "" {
        out = strftime.Format(c.String("format"), result)
      }
      fmt.Println(out)
    }
  }

  app.Run(os.Args)
}
