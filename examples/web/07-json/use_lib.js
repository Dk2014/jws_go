// Generated by CoffeeScript 1.9.3
(function() {
  var JsonEditor, JsonModByUrl, reqJson;

  reqJson = function(num, cb) {
    return $.ajax({
      type: "GET",
      url: "../api/v1/get",
      data: "num=" + num,
      dataType: "text",
      success: cb
    });
  };

  JsonModByUrl = React.createClass({
    getInitialState: function() {
      return {
        editor: null
      };
    },
    componentDidMount: function() {
      var change_cb, container, editor_imp, options;
      container = document.getElementById(this.props.name + "jsoneditor");
      change_cb = this.onchange.bind(this);
      options = {
        "mode": "tree",
        "search": true,
        change: change_cb
      };
      editor_imp = new JSONEditor(container, options);
      this.getData(function(data) {
        editor_imp.set(JSON.parse(data));
        change_cb(data);
      });
      this.setState({
        editor: editor_imp
      });
    },
    getData: function(cb) {
      reqJson(22, cb);
    },
    onchange: function() {
      var data;
      data = this.state.editor.getText();
      return this.props.onchange(data);
    },
    render: function() {
      return React.createElement("div", {
        "id": this.props.name + "jsoneditor"
      });
    }
  });

  JsonEditor = React.createClass({
    getInitialState: function() {
      return {
        json: "{}"
      };
    },
    componentDidMount: function() {},
    onchange: function(data) {
      console.log("onchange");
      console.log(data);
      this.setState({
        json: data
      });
      return console.log(this.state);
    },
    render: function() {
      return React.createElement("div", null, React.createElement("input", {
        "type": "text",
        "value": this.state.json
      }), React.createElement(JsonModByUrl, {
        "id": name,
        "key": name,
        "name": name,
        "onchange": this.onchange
      }));
    }
  });

  React.render(React.createElement(JsonEditor, null), document.getElementById('example4'));

}).call(this);
