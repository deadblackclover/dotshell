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

switch ($Task) {
    "fmt" { Fmt }
    "build" { Build }
    "run" { Run }
    Default { Write-Host "Unknown task: $Task" }
}
