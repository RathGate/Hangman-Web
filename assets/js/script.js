function resetGame() {
    var url = "../../main.go"
    $.ajax({
        type: "POST",
        url: url,
        data: { "reset": true },
        success: function(data) {
            location.reload()
        }
    })
}