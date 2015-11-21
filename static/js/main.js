var isEventOverDiv = function(x, y) {

  var external_events = $('#external-events');
  var offset = external_events.offset();
  offset.right = external_events.width() + offset.left;
  offset.bottom = external_events.height() + offset.top;

  // Compare
  if (x >= offset.left && y >= offset.top && x <= offset.right && y <= offset.bottom) {
    return true;
  }
  return false;

}
$(function() {
  $('#calendar').fullCalendar({
    header: {
      left: 'title',
      center: '',
      right: 'today prev,next month,agendaWeek,agendaDay'
    },
    defaultView: 'agendaWeek',
    droppable: true,
    editable: true,
    drop: function(date, jsEvent, ui) {
      this.remove();
    },
    dragRevertDuration: 0,
    eventDragStop: function(event, jsEvent, ui, view) {

      if (isEventOverDiv(jsEvent.clientX, jsEvent.clientY)) {
        $('#calendar').fullCalendar('removeEvents', event._id);
        var el = $("<div class='draggable' data-event='{'title':'my event'}'>").appendTo('#external-events-listing').text(event.title);
        el.draggable({
          zIndex: 999,
          revert: true,
          revertDuration: 0
        });
        el.data('event', {
          title: event.title,
          id: event.id,
          stick: true
        });
      }
    }
  });
  $(".draggable").draggable({
    revert: true,
    revertDuration: 0
  });
});
