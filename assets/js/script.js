window.onload = () => {
    radioListener()

    $( "#restart-btn" ).click(function() {
        $.ajax({
            type: "POST",
            url: "/hangman",
            data: { "difficulty": "current" },
        });
        window.location.href = "/hangman"
      });
}





function radioListener() {
    radio = document.querySelectorAll(`[name="difficulty"]`)
    radio.forEach(input => input.addEventListener('change', function() {
        $.ajax({
            type: "POST",
            url: "/",
            data: { "requestedData": input.value },
            success: function(data) {
                updateScoreboard(data)
            }
        })
    }))
}

function updateScoreboard(data) {
    if (data == "404") {
        return
    } else {
        data = JSON.parse(data)
        document.getElementById("diffname").innerHTML = data.Name
        document.getElementById("winstreak").innerHTML = data.CurrentWinStreak
        document.getElementById("wincount").innerHTML = data.WinCount
        document.getElementById("losecount").innerHTML = data.LoseCount
    }
}
