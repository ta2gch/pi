download_1(){
    from="${1}"
    dist="${2}"
    if which wget > /dev/null ; then
	wget -O "${dist}" "${from}"
    elif which curl > /dev/null ; then
	if [ "${dist}" = "-" ] ; then
	    curl -L -O "${from}"
	else
	    curl -L -o "${dist}" -O "${from}"
	fi
    fi
}

download(){
    from="${1}"
    dist="${2}"
    case "${from}" in
	https://*) download_1 "${from}" "${dist}" ;;
	http://*) download_1 "${from}" "${dist}" ;;
	*) cp "${from}" "${dist}" ;;
    esac
}
