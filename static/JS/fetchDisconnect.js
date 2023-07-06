const onClickDisconnect = () => {
    fetch("/logout")
    window.location.href = "/"
}