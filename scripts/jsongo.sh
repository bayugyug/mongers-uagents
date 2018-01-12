#!/bin/bash 







function _loadIt(){
    local fn=${1:2}
    echo loading $fn

    echo " 
    package main

    import () 


    var UA_${fn} = []string{
    $(_dump $fn)
}
    "
}


function _dump(){
    cat $1 | awk '{printf("\"%s\",\n",$0)}'
}

for fn in $(find . -type f -name '*.json' 2>/dev/null)
{
    _loadIt $fn > $fn.go
}



