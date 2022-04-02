# debugdefsec

Debug for [defsec](https://github.com/aquasecurity/defsec).

# Usage

```shell
$ debugdefsec dockerfile testdata/Dockerfile | jq
{
  "Stages": {
    "hub.example.com/ubuntu:20.04": [
      {
        "Cmd": "from",
        "SubCmd": "",
        "Flags": [],
        "Value": [
          "hub.example.com/ubuntu:20.04"
        ],
        "Original": "FROM hub.example.com/ubuntu:20.04",
        "JSON": false,
        "Stage": 0,
        "Path": "testdata/Dockerfile",
        "StartLine": 2,
        "EndLine": 2
      },
      {
        "Cmd": "run",
        "SubCmd": "",
        "Flags": [],
        "Value": [
          "apt-get update && apt-get install curl"
        ],
        "Original": "RUN apt-get update && apt-get install curl",
        "JSON": false,
        "Stage": 0,
        "Path": "testdata/Dockerfile",
        "StartLine": 4,
        "EndLine": 4
      },
      {
        "Cmd": "expose",
        "SubCmd": "",
        "Flags": [],
        "Value": [
          "22"
        ],
        "Original": "EXPOSE 22",
        "JSON": false,
        "Stage": 0,
        "Path": "testdata/Dockerfile",
        "StartLine": 5,
        "EndLine": 5
      }
    ]
  }
}
```
