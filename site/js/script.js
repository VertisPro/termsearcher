function autocomplet() {
	var min_length = 4; // min caracters to display the autocomplete
	var keyword = $('#searchstring').val();
	if (keyword.length >= min_length) {
		$.ajax({
			url: '/search/',
			type: 'POST',
			data: {searchstring:keyword},
			success:function(data){
				$('#search_results').show();
				$('#search_results').html(data);
			}
		});
	} else {
		$('#search_results').hide();
	}
}
 
// set_item : this function will be executed when we select an item
function set_item(item) {
	// change input value
	$('#searchstring').val(item);
	// hide proposition list
	$('#search_results').hide();
}


// attach ready event
$(document).ready(function () {
	//checkbox
	var $checkbox = $('.ui.toggle.checkbox'), handler;
	// event handlers
	handler = { };
	// activate the checkbox
	$checkbox.checkbox();

	//dropdown
	$('.dropdown').dropdown({
    });

	var $searchtype = $('.dropdown').dropdown('get text');

	// initializes with default endpoint /search/{query}
  $('.ui.search')
    .search({
      apiSettings: {
        url: '/search/{query}'
      },
      type: 'category'
    })
  ;

});