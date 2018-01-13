## mongers-uagents


- [x] Simple golang script that will do a list of best known user agents string of mobile devices


- [x] Output the details of the user agent string in JSON format



## Compile

```sh


```

## Usage

```sh

$ ./mongers-uagents


Usage: ./mongers-uagents [options]

     Options are:


  -a string
        Show the user agent details (short form)
  -agent string
        Show the user agent details
  -l string
        Show all the user agent list (short form)
  -list string
        Show all the user agent list
  -s    Show the user agent main list (short form)
  -show
        Show the user agent main list


```

## Show the user agent details

- [x] Show it in JSON format

``` sh

./mongers-uagents -a "Mozilla/5.0 (Linux; U; Android 7.1; en-US; Pixel Build/NDE63H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 UCBrowser/11.3.2.960 U3/0.8.0 Mobile Safari/534.30",
{
        "UserAgent": {
                "Family": "UC Browser",
                "Major": "11",
                "Minor": "3",
                "Patch": "2"
        },
        "Os": {
                "Family": "Android",
                "Major": "7",
                "Minor": "1",
                "Patch": "",
                "PatchMinor": ""
        },
        "Device": {
                "Family": "Pixel",
                "Brand": "Generic_Android",
                "Model": "Pixel"
        }
}



./mongers-uagents -a "Mozilla/5.0 (iPhone; U; CPU iPhone OS 10_0 like Mac OS X; en-us) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/7A341 Safari/602.1",
{
        "UserAgent": {
                "Family": "Mobile Safari",
                "Major": "10",
                "Minor": "0",
                "Patch": ""
        },
        "Os": {
                "Family": "iOS",
                "Major": "10",
                "Minor": "0",
                "Patch": "",
                "PatchMinor": ""
        },
        "Device": {
                "Family": "iPhone",
                "Brand": "Apple",
                "Model": "iPhone"
        }
}



```


## Show the user agent main list

- [x] Show it in JSON format

``` sh
./mongers-uagents -s

{
        "agent": "UserAgents",
        "total": 13,
        "data": [
                "Android",
                "LinuxOs",
                "UnixOs",
                "Ios",
                "Mac",
                "MacOs",
                "WindowsOs",
                "Blackberry",
                "ChromeOs",
                "WindowsPhone",
                "SymbianOs",
                "FireOs",
                "WindowsMobile"
        ]
}


```

## Docker Binary

- [x] In order to  use it as dockerize binary


``` sh


    sudo  sysctl -w net.ipv4.ip_forward=1

    sudo  docker run --rm  bayugyug/mongers-uagents

```


### License

[MIT](https://bayugyug.mit-license.org/)
