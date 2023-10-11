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
                img.src = video.ImageFilePath

                const cardText = videoItem.querySelector(".card-text")
                cardText.textContent = video.Description
                
                const durationBlock = videoItem.querySelector(".video-duration")
                durationBlock.textContent = getDurationText(video.DurationSeconds)

                videosElement.appendChild(videoItem)
            }
        }).catch(error => {
            console.error(error)
        })

        let videosLink = document.getElementById("videos-link")
        videosLink.addEventListener("click", onVideosLinkClick)

        let settingsLink = document.getElementById("settings-link")
        settingsLink.addEventListener("click", onSettingsLinkClick)
})



    function refreshImages() { 
        const imageElements = document.querySelectorAll('#videos-list img');

        // Iterate through each image element and update its source attribute
        imageElements.forEach((imgElement) => {
            // Iterates over all images (using the for.each) and generates a random image URL using picsum.photos.  https://www.w3schools.com/js/js_callback.asp
            const randomImageUrl = `https://picsum.photos/${imgElement.width}/${imgElement.height}?random=${Math.random()}`; 

            // math random used to generate new pic every time - https://www.w3schools.com/js/js_random.asp
            
            // Set the new image source. .src is short for source from url https://www.w3schools.com/jsref/prop_script_src.asp
            imgElement.src = randomImageUrl;
        });
    }

    
    refreshImages();  

    function onVideosLinkClick() { 
        let videosListContainer = document.getElementById("videos-list-container") //document.get used to get HTML element from the video list container 
        videosListContainer.classList.remove("d-none") // d none used to hide element 

        let settingsListContainer = document.getElementById("settings-list-container") 
        settingsListContainer.classList.add("d-none")

        
        refreshImages(); // refreshes images when button is pressed 
    }

    function onSettingsLinkClick() {
        let videosListContainer = document.getElementById("videos-list-container")
        videosListContainer.classList.add("d-none")

        let settingsListContainer = document.getElementById("settings-list-container")
        settingsListContainer.classList.remove("d-none")
    }

    
    const refreshButton = document.getElementById("main-call-button"); // main-call-button is the button identifier 

    
    refreshButton.addEventListener("click", refreshImages); // button that calls the refreshImages function

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

