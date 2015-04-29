function loadCodeSnippet() {
   var snippetBox = document.getElementById("codeSnippetArea");

   // selectionStart and selectionEnd are mutable by the user, so
   // we'll use our own property to keep track of where the
   // user is.
   // __ is the prefix used for the property since we're technically not 
   // allowed to extend DOM objects in an adhoc fashion (but everyone does it 
   // anyways).
   snippetBox.__cursorPos = 0;

   // grab a raw-text code snippet
   snippetBox.value = "Please wait while we fetch your code snippet ...";
   httpGetAsync("/snippet", function(response) {
      snippetBox.value = response;
      snippetBox.focus();
      snippetBox.scrollTop = 0;
   });

   snippetBox.onkeypress = function() {
      snippetBox.__cursorPos = snippetBox.__cursorPos + 1;
      snippetBox.setSelectionRange(0, snippetBox.__cursorPos);
   };
}

function httpGetAsync(url, func) {
    var xmlHttp = null;

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", url, true);
    xmlHttp.onload = function(e) {func(xmlHttp.responseText)};
    xmlHttp.send(null);
}
