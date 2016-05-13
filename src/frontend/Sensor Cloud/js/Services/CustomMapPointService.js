//custom map point service used for sharing coordidata 
//between two or more controllers
angular
.module('sensorDataModule')
.service('CustomMapPointService', CustomMapPointService);

function CustomMapPointService(){
	var object = this;
	var lat = "";
	var lng = "";

	obj.getLat = function(){
		return lat;
	}
	obj.setLat = function(latInput){
		lat = latInput;
	}
	
	obj.getLng = function(){
		return lng;
	}
	obj.setLng = function(lngInput){
		lng = lngInput;
	}
	return object;
}