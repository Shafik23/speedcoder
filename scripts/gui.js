function loadCodeSnippet() {
   var snippetBox = document.getElementById("codeSnippetArea");

   // selectionStart and selectionEnd are mutable by the user, so
   // we'll use our own property to keep track of where the
   // user is.
   snippetBox.cursorPos = 0;

   // grab a raw-text code snippet
   snippetBox.value = httpGet("/snippet");

   snippetBox.onkeypress = function() {
      snippetBox.cursorPos = snippetBox.cursorPos + 1;
      snippetBox.setSelectionRange(0, snippetBox.cursorPos);

      console.log("cursorPos is : " + snippetBox.cursorPos);
   }
}

function httpGet(url) {
    var xmlHttp = null;

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", url, false);
    xmlHttp.send(null);
    return xmlHttp.responseText;
}
