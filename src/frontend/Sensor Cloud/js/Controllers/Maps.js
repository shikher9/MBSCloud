angular
.module('sensorDataModule')
.controller('GoogleMapLocController', GoogleMapLocController);
GoogleMapLocController.$inject = ['$scope', 'CustomMapPointService'];
function GoogleMapLocController($scope, CustomMapPointService){
	$scope.setCurrentCoordinates = function(lat,lng){
		
		// setting coordinates for sharing
		$scope.lat = CustomMapPointService.setLat(lat);
		$scope.lng = CustomMapPointService.setLng(lng);	
	}
}

if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(successFunction, errorFunction);
}

function successFunction(position) {
    var mapOptions = {
        center: new google.maps.LatLng(position.lat, position.lng),
        zoom: 6,
        mapTypeId: google.maps.MapTypeId.ROADMAP
    };
    var infoWindow = new google.maps.InfoWindow();
    var latlngbounds = new google.maps.LatLngBounds();
    var map = new google.maps.Map(document.getElementById("dvMap"), mapOptions);
    google.maps.event.addListener(map, 'click', function (e) {
        document.getElementById('lat').value = e.latLng.lat();
        document.getElementById('lng').value = e.latLng.lng();
        var lat =e.latLng.lat();
        var lng = e.latLng.lng();
     
    });

}

function errorFunction(){
    alert("Geocoder failed");
}


