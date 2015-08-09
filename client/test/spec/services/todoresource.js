'use strict';

describe('Service: TodoResource', function () {

  // load the service's module
  beforeEach(module('todoApp'));

  // instantiate service
  var TodoResource;
  beforeEach(inject(function (_TodoResource_) {
    TodoResource = _TodoResource_;
  }));

  it('should do something', function () {
    expect(!!TodoResource).toBe(true);
  });

});
