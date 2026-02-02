param (
    [Parameter(Mandatory = $false)]
    [string]$Task = "build"
)

function Fmt {
    gofmt -s -w ./.
}

function Build {
    go build -o ./build/ ./cmd/dotshell
}

function Run {
    go run ./cmd/dotshell
}

function RunPublic {
    $env:HOST="0.0.0.0"
    Run
}

switch ($Task) {
    "fmt" { Fmt }
    "build" { Build }
    "run" { Run }
    "run-public" { RunPublic }
    Default { Write-Host "Unknown task: $Task" }
}
