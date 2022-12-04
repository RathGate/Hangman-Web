window.onload = () => {
    radioListener()
    document.getElementById("difficulty-form").onload = defaultCheck()
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
        document.getElementById("winstreak").innerHtml = data.CurrentWinStreak
        document.getElementById("wincount").innerHtml = data.WinCount
        document.getElementById("losecount").innerHtml = data.LoseCount
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