// Generated by CoffeeScript 1.9.3
(function() {
  var JsonMod, js, json, mkJsonMod;

  JsonMod = React.createClass({
    getInitialState: function() {
      return {};
    },
    componentDidMount: function() {
      var container, editor, options;
      container = document.getElementById(this.props.name + "jsoneditor");
      options = {
        "mode": "tree",
        "search": true
      };
      editor = new JSONEditor(container, options);
      editor.set(JSON.parse(this.props.data));
    },
    render: function() {
      return React.createElement("div", {
        "id": this.props.name + "jsoneditor"
      });
    }
  });

  mkJsonMod = function(data, name) {
    return React.createElement("div", null, React.createElement(JsonMod, {
      "id": name,
      "key": name,
      "data": data,
      "name": name
    }));
  };

  json = {
    "Array": [1, 2, 3],
    "Boolean": true,
    "Null": null,
    "Number": 123,
    "Object": {
      "a": "b",
      "c": "d"
    },
    "String": "Hello World"
  };

  js = mkJsonMod(JSON.stringify(json), "test");

  React.render(js, document.getElementById('example3'));

}).call(this);
