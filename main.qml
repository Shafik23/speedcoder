import QtQuick 2.0
import QtQuick.Controls 1.1
import QtQuick.Controls.Styles 1.1
import GoExtensions 1.0

Rectangle {
   width: 1024
   height: 768
   color: "black"

   TextArea {
      id: textArea
      width: parent.width * 0.9
      height: parent.height *0.9
      anchors.centerIn: parent
      focus: true

      property var typedIndex: 0

      backgroundVisible: true
      selectByMouse: false
      readOnly: true

      style: TextAreaStyle {
         textColor: "darkgrey"
         selectionColor: "black"
         selectedTextColor: "lightgreen"
         backgroundColor: "black"
      }

      font.family: "Monospace"
      font.pointSize: 16
      font.bold: false
      textMargin: 15

      text: snippet.code

      Snippet {
         id: snippet
      }

      function keyStrokeMatches(input, snippet_code_char) {
         // if either key or snippet char is 10 and the other is 13, consider it a match
         // this makes carriage returns match newlines
         if (input.charCodeAt(0) === 10 && snippet_code_char.charCodeAt(0) === 13 ||
             input.charCodeAt(0) === 13 && snippet_code_char.charCodeAt(0) === 10) {
            return true;
         }

         // if the input is the Esc key, match it so the user can just skip ahead
         if (input.charCodeAt(0) === 27) return true;

         return input === snippet_code_char;
      }

      Keys.onPressed: {
         if (keyStrokeMatches(event.text, textArea.text[typedIndex])) {
            typedIndex = typedIndex + 1
         }

         textArea.select(0, typedIndex)
      }

      /*
       Gopher {

          MouseArea {
             anchors.fill: parent

             property real startX
             property real startR

             onPressed: {
                startX = mouse.x
                startR = gopher.rotation
                anim.running = false
             }
             onReleased: {
                anim.from = gopher.rotation + 360
                anim.to = gopher.rotation
                anim.running = true
             }
             onPositionChanged: {
                gopher.rotation = (36000 + (startR - (mouse.x - startX))) % 360
             }
          }

       }
       */
   }
}
