document.addEventListener("DOMContentLoaded", function(){
    fetch("/videos")
        .then(response => response.json())
        .then(data => {
            const videosElement = document.getElementById('videos-list')
            for (var i = 0; i < data.length; i++) {
                var video = data[i]

                const videoTemplate = document.getElementById("video-item")
                const videoItem = videoTemplate.content.cloneNode(true)

                const img = videoItem.querySelector("img")
                img.src =  "/static/img/" + video.Image

                const cardText = videoItem.querySelector(".card-text")
                cardText.textContent = video.Text
                
                const durationBlock = videoItem.querySelector(".video-duration")
                durationBlock.textContent = getDurationText(video.DurationSeconds)

                videosElement.appendChild(videoItem)
            }
        }).catch(error => {
            console.error(error)
        })
})

function getDurationText(durationSeconds) {
    const secondsInMinute = 60
    let secondText = "second"

    if (durationSeconds == 0) {
        return ""
    }
    if (durationSeconds == 1) {
        return joinText([durationSeconds, secondText])
    }

    if (durationSeconds < secondsInMinute) {
        secondText += "s"
        return joinText([durationSeconds, secondText])
    }

    let minuteText = "minute"
    const durationMinutes = Math.floor(durationSeconds / 60)
    if (durationMinutes == 1) {
        return joinText([durationMinutes, minuteText])
    }

    minuteText +="s"
    return joinText([durationMinutes, minuteText])
}

function joinText(textItems) {
    return textItems.join(" ");
}