window.onload = () => {
    radioListener()
    defaultCheck()
}

function radioListener() {
    radio = document.querySelectorAll(`[name="difficulty"]`)
    radio.forEach(input => input.addEventListener('change', function() {
        $.ajax({
            type: "POST",
            url: "/",
            data: { "getDifficultyData": input.value },
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

function defaultCheck() {
    if ($("#diffname").length) {
        let content = document.getElementById("diffname").innerText
        document.getElementById(`diff-${content}`).checked = true
        return
    } 
    document.getElementById("diff-medium").checked = true;
}
