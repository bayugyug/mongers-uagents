## gohttpdummy


- [x] This is a simple golang script that will do a list of best known user agents of mobile devices


- [x] Output the details of the user agent



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

./mongers-uagents -a  "Mozilla/5.0 (Linux; Android 7.1.1; SAMSUNG GT-I9505 Build/NOF27B) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/5.2 Chrome/51.0.2704.106 Mobile Safari/537.36"
{
        "UserAgent": {
                "Family": "Samsung Internet",
                "Major": "5",
                "Minor": "2",
                "Patch": ""
        },
        "Os": {
                "Family": "Android",
                "Major": "7",
                "Minor": "1",
                "Patch": "1",
                "PatchMinor": ""
        },
        "Device": {
                "Family": "Samsung GT-I9505",
                "Brand": "Samsung",
                "Model": "GT-I9505"
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
