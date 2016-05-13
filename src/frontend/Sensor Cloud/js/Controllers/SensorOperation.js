angular
.module('sensorDataModule')
.controller('SensorOpController', SensorOpController);
SensorOpController.$inject = ['$scope', 'CustomMapPointService'];
function SensorOpController($scope, CustomMapPointService){
	$scope.getCurrentCoordinates = function(){
		$scope.lat = CustomMapPointService.getLat();
		$scope.lng = CustomMapPointService.getLng();
	}
}