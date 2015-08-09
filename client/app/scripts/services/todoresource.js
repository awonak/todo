'use strict';

/*
  Todo API
  GET: /api/tasks
  POST: /api/tasks
  GET: /api/tasks/:id
  POST: /api/tasks/:id
  DELETE: /api/tasks/:id
 */
angular.module('todoApp')
  .factory('Tasks', function ($resource, API_SERVER) {
    var urlBase = API_SERVER + 'tasks';

    // Public API here
    return $resource(urlBase + '/:id', {id:'@id'});

  });