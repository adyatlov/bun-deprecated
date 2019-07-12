package errors

import (
  "bufio"
  "fmt"
  "regexp"
  "github.com/adyatlov/bun"
)

// http://erlang.org/doc/apps/kernel/logger_chapter.html#log_level
var errorsRegex = "\\[error\\]|\\[critical\\]|\\[alert\\]|\\[emergency\\]"

func init() {
  builder := bun.CheckBuilder{
    Name: "net-errors-checker",
    Description: "Identify errors in dcos-net logs",
    CollectFromMasters:      collect,
    CollectFromAgents:       collect,
    CollectFromPublicAgents: collect,
    Aggregate:               aggregate,
  }
  check := builder.Build()
  bun.RegisterCheck(check)
}

func collect(host bun.Host) (ok bool, details interface{}, err error) {
  errorMatcher := regexp.MustCompile(errorsRegex)
  keys := 0

  file, err := host.OpenFile("net")

  if err != nil {
    ok = false
    errMsg := fmt.Sprintf("Cannot open net file %s", err)
    fmt.Println(errMsg)
    return
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if errorMatcher.MatchString(line) {
      keys++
    }
  }

  details = keys
  ok = true
  return
}

func aggregate(c *bun.Check, b bun.CheckBuilder) {
  details := []string{}
  ok := true

  c.Summary = fmt.Sprintf("There are no networking errors on dcos-net")

  for _, r := range b.OKs {
    v := r.Details.(int)
    if v != 0 {
      ok = false
      details = append(details, fmt.Sprintf("%d log lines at %v %v with level error or bellow",
        v, r.Host.Type, r.Host.IP))
    }
  }

  if ok {
    c.OKs = details
    c.Summary = "dcos-net logs are OK."
  } else {
    c.Problems = details
    c.Summary = fmt.Sprintf("There are errors on dcos-net logs on %d out of %d nodes.", len(details), len(b.OKs))
  }
}

