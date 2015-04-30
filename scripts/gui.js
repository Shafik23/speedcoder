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
   setGoButtonState(false);
   snippetBox.value = "Please wait while we fetch your code snippet ...";
   httpGetAsync("/snippet", function(response) {
      snippetBox.value = response;
      snippetBox.focus();
      snippetBox.scrollTop = 0;
      setGoButtonState(true);
   });

   // convenience function
   var advance = function(n) {snippetBox.__cursorPos = snippetBox.__cursorPos + n;} ;

   snippetBox.onkeypress = function(event) {
      var currentCharCode = snippetBox.value.charCodeAt(snippetBox.__cursorPos);

      // only advance the cursor if the user types the correct character
      if (currentCharCode === event.charCode) {
         advance(1);
      }

      // line-feeds and carriage returns are considered the same
      if (currentCharCode === 10 && event.charCode === 13 ||
            currentCharCode === 13 && event.charCode === 10) {
         advance(1);
      }

      snippetBox.onfocus();
      return false;
   };

   // keydown is necessary (vs keypress) in order to detect special keys
   // like ESC
   snippetBox.onkeydown = function(event) {
      // if the user presses the ESC key, they can "bypass" chunks of displayed
      // code quickly. This comes in handy when going through boring code comments,
      // for example
      if (event.keyCode === 27) {
         advance(5);
         snippetBox.onfocus();
         return false;
      }
   }

   snippetBox.onfocus = function() {
      snippetBox.setSelectionRange(0, snippetBox.__cursorPos);
   }
}

function setGoButtonState(flag) {
   document.getElementById("GoButton").disabled = !flag;
}

/**
 * Performs an async XMLHttpRequest, calling func on the 
 * response once it's available.
 */
function httpGetAsync(url, func) {
    var xmlHttp = null;

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", url, true);
    xmlHttp.onload = function(e) {func(xmlHttp.responseText)};
    xmlHttp.send(null);
}
