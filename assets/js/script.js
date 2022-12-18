function loadJose() {
    if (!$("#jose").length || attempts == 10) {
        return
    }

    let divs = document.querySelectorAll('[attempts]')

    for (let i = 0; i < 10 - attempts; i++) {
        divs[i].classList.add("visible")
    }
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
        document.getElementById("maxwinstreak").innerHTML = data.MaxWinStreak
        document.getElementById("wincount").innerHTML = data.WinCount
        document.getElementById("losecount").innerHTML = data.LoseCount
        getRatio()
    }
}

function getRatio() {
    console.log("bonjour ?????")
    if (!$(".scoreboard-data").length) {
        return
    }
    
    let wins = parseInt($("#wincount").text())
    let total = wins + parseInt($("#losecount").text())
    let ratioDiv = document.getElementById("ratio")

    if (total == 0) {
        ratioDiv.innerText = 0; ratioDiv.style.color = "white"
        return
    }
        
    var winrate = Math.round(wins / total * 100);
    
    document.getElementById("ratio").innerText = winrate
    document.getElementById("ratio").style.color = getRGB(winrate)
}

function checkRGB(value) {
    if (value < 0) {
        return 0
    } else if (value > 255) {
        return 255
    }
    return value
}

function getRGB(perc) {
    if (perc < 50) {
        // If 40
        var calc = checkRGB(Math.round(255 * ((50 - perc) * 2 / 100)))
        return (`rgb(255, ${calc}, 0)`)
    } else if (perc > 50) {
        var calc = checkRGB(Math.round(255 - (255 * ((perc - 50) * 2 / 100))))
        return (`rgb(${calc}, 255, 0)`)
    }
    return "rgb(255, 255, 0)"
}

function main() {
    loadJose()
    radioListener()
    getRatio()
}

main()