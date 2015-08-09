'use strict';

angular.module('todoApp')
  .controller('MainCtrl', function ($scope, Tasks) {

    $scope.todos = [];

    var tasks = Tasks.query(function() {
      $scope.todos = tasks;
    });

    $scope.addTodo = function () {
      var task = new Tasks();
      task.title = $scope.todo;
      task.$save();

      $scope.todos.push(task);
      $scope.todo = '';
    };

    $scope.removeTodo = function (index) {
      Tasks.delete({id: $scope.todos[index].id});
      $scope.todos.splice(index, 1);
    };

  });
