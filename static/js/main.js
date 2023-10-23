document.addEventListener("DOMContentLoaded", function () {
  fetch("/videos")
    .then((response) => response.json())
    .then((data) => {
      const videosElement = document.getElementById("videos-list");
      for (var i = 0; i < data.length; i++) {
        var video = data[i];

        const videoTemplate = document.getElementById("video-item");
        const videoItem = videoTemplate.content.cloneNode(true);

        const img = videoItem.querySelector("img");
        img.src = video.ImageFilePath;

        const cardText = videoItem.querySelector(".card-text");
        cardText.textContent = video.Description;

        const durationBlock = videoItem.querySelector(".video-duration");
        durationBlock.textContent = getDurationText(video.DurationSeconds);

        videosElement.appendChild(videoItem);
      }
    }).catch((error) => {
      console.error(error);
    });

  
    function FlashMessageSuccess(successBtn, messagePlaceholder) {
      // Flash message success 
  const button = document.getElementById(successBtn);
  const successAlertPlaceholder = document.getElementById(messagePlaceholder);

  button.addEventListener("click", function () {
    // Success alert
    const successAlert = document.createElement("div");
    successAlert.classList.add("alert", "alert-success");
    successAlert.textContent = "Successful!";

    successAlertPlaceholder.innerHTML = "";
    successAlertPlaceholder.appendChild(successAlert);

    // Timer set to 5 seconds
    setTimeout(function () {
      successAlertPlaceholder.innerHTML = "";
    }, 5000);
  });
    }

    FlashMessageSuccess("main-call-button", "success-alert-placeholder")

    FlashMessageSuccess("save-settings-button", "confirmationMessage")

    

  let videosLink = document.getElementById("videos-link");
  videosLink.addEventListener("click", onVideosLinkClick);

  let settingsLink = document.getElementById("settings-link");
  settingsLink.addEventListener("click", onSettingsLinkClick);

  let saveSettingsButton = document.getElementById("save-settings-button");
  saveSettingsButton.addEventListener("click", writeSettings);

  let deleteSettingsButton = document.getElementById("delete-settings-button");
  deleteSettingsButton.addEventListener("click", deleteSettings);
});

function onVideosLinkClick() {
  let videosListContainer = document.getElementById("videos-list-container");
  videosListContainer.classList.remove("d-none");

  let settingsListContainer = document.getElementById(
    "settings-list-container"
  );
  settingsListContainer.classList.add("d-none");
}

function onSettingsLinkClick() {
  let videosListContainer = document.getElementById("videos-list-container");
  videosListContainer.classList.add("d-none");

  let settingsListContainer = document.getElementById(
    "settings-list-container"
  );
  settingsListContainer.classList.remove("d-none");

  readSettings();
}

function getDurationText(durationSeconds) {
  const secondsInMinute = 60;
  let secondText = "second";

  if (durationSeconds == 0) {
    return "";
  }
  if (durationSeconds == 1) {
    return joinText([durationSeconds, secondText]);
  }

  if (durationSeconds < secondsInMinute) {
    secondText += "s";
    return joinText([durationSeconds, secondText]);
  }

  let minuteText = "minute";
  const durationMinutes = Math.floor(durationSeconds / 60);
  if (durationMinutes == 1) {
    return joinText([durationMinutes, minuteText]);
  }

  minuteText += "s";
  return joinText([durationMinutes, minuteText]);
}

function joinText(textItems) {
  return textItems.join(" ");
}

function readSettings() {
  fetch("/settings")
    .then((response) => response.json())
    .then((data) => {
      let aboutTitleInput = document.getElementById("aboutTitleInput");
      aboutTitleInput.value = data.AboutTitle;

      let aboutTextInput = document.getElementById("aboutTextInput");
      aboutTextInput.value = data.AboutText;

      let videosCountOnMainPageInput = document.getElementById(
        "videosCountOnMainPageInput"
      );
      videosCountOnMainPageInput.value = data.VideosCountOnMainPage;
    })
    .catch((error) => {
      console.error(error);
    });
}

function writeSettings() {
  const form = document.getElementById("setting-form");
  const data = new FormData(form);
  const plainFormData = Object.fromEntries(data.entries());

  const formDataJsonString = JSON.stringify(plainFormData);

  fetch("/persist-settings", {
    body: formDataJsonString,
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    }
  }).then((response) => response.json())
  .then((data) => {
      console.log(data);
  })
  .catch((error) => {
      console.error(error);
  });
}

    function refreshImages() {
        
        const imageElements = document.querySelectorAll('#videos-list img');

        // Iterate through each image element and update its source attribute
        imageElements.forEach((imgElement) => {
            // Generate a random image URL using picsum.photos
            const randomImageUrl = `https://picsum.photos/${imgElement.width}/${imgElement.height}?random=${Math.random()}`;
            
            
            imgElement.src = randomImageUrl;
        });
    }

    // Call the refreshImages function to refresh the images initially
    refreshImages();

    function onVideosLinkClick() {
        let videosListContainer = document.getElementById("videos-list-container")
        videosListContainer.classList.remove("d-none")

        let settingsListContainer = document.getElementById("settings-list-container")
        settingsListContainer.classList.add("d-none")

        
        refreshImages();
    }

    function onSettingsLinkClick() {
        let videosListContainer = document.getElementById("videos-list-container")
        videosListContainer.classList.add("d-none")

        let settingsListContainer = document.getElementById("settings-list-container")
        settingsListContainer.classList.remove("d-none")
    }

    // Get a reference to the "Refresh Images" button by its ID
    const refreshButton = document.getElementById("main-call-button");

    // Add a click event listener to the button that calls the refreshImages function
    refreshButton.addEventListener("click", refreshImages);


function deleteSettings() {
  fetch("/delete-settings", {
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    }
  }).then((response) => response.json())
  .then((data) => {
      console.log(data);
  })
  .catch((error) => {
      console.error(error);
  });
}
