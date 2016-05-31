function leap() {
    local LEAP
    LEAP="$GOPATH/bin/leap"
 
    if [ $# -lt 1 ]; then
	$LEAP
    elif [ $1 = "add" ] || [ $1 = "rm" ] || [ $1 = "list" ] || [ $1 = "help" ]; then
	$LEAP $@
    else
        cd $($LEAP "$@")
    fi
}
