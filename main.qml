import QtQuick 2.0
import QtQuick.Controls 1.1
import GoExtensions 1.0

Rectangle {
   width: 1024
   height: 768
   color: "black"

   TextArea {
      width: parent.width * 0.9
      height: parent.height *0.9
      anchors.centerIn: parent

      backgroundVisible: false
      readOnly: true

      font.family: "Helvetica"
      font.pointSize: 20
      textColor: "green"
      text: snippet.code

      Snippet {
         id: snippet
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
