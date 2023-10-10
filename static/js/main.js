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
    })
    .catch((error) => {
      console.error(error);
    });

  let videosLink = document.getElementById("videos-link");
  videosLink.addEventListener("click", onVideosLinkClick);

  let settingsLink = document.getElementById("settings-link");
  settingsLink.addEventListener("click", onSettingsLinkClick);

  let saveSettingsButton = document.getElementById("save-settings-button");
  saveSettingsButton.addEventListener("click", writeSettings);
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
  const minutes = Math.floor((durationSeconds % 3600) / 60);
  const remainingSeconds = Math.floor(durationSeconds % 60);

  const minuteStr = minutes > 0 ? minutes + "m " : "";
  const secondStr = remainingSeconds != 0 ? remainingSeconds + "s" : "";

  return minuteStr + secondStr;
  
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
  plainFormData.VideosCountOnMainPage = -1
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
