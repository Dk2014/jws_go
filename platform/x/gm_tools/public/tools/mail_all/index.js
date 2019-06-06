// Generated by CoffeeScript 1.12.7
(function() {
  var AccountKeyInput, Api, BatchMailResult, Button, Icon, Input, MailAll, MailBatch, MailList, MailOne, NewMailAllSendModal, React, Upload, antd, boot, maillist;

  NewMailAllSendModal = require('./new_mail_to_all');

  maillist = require('./mail_list');

  MailList = maillist.MailList;

  BatchMailResult = maillist.BatchMailResult;

  React = require('react');

  AccountKeyInput = require('../../common/account_input');

  boot = require('react-bootstrap');

  antd = require('antd');

  Upload = antd.Upload;

  Icon = antd.Icon;

  Button = antd.Button;

  Input = antd.Input;

  Api = require('../api/api_ajax');

  MailAll = React.createClass({displayName: "MailAll",
    getInitialState: function() {
      return {
        select_server: this.props.curr_server,
        player_to_send: "0:0:1331"
      };
    },
    handleServerChange: function(data) {
      console.log(data);
      this.setState({
        select_server: data
      });
      return setTimeout(this.OnSend, 100);
    },
    handleUserChange: function(data) {
      console.log(data);
      if (data === "") {
        data = "请输入玩家Id";
      }
      return this.setState({
        player_to_send: data
      });
    },
    OnSend: function() {
      console.log("SendOver");
      return this.refs.mail_list_all.Refersh(1);
    },
    getAllMailName: function() {
      var res;
      if (this.state.select_server != null) {
        res = this.state.select_server.serverName.split(':');
        return "all:" + res[1];
      } else {
        return "all:10";
      }
    },
    getServerName: function() {
      if (this.state.select_server != null) {
        return this.state.select_server.serverName;
      } else {
        return "";
      }
    },
    render: function() {
      return React.createElement("div", null, React.createElement(NewMailAllSendModal, {
        "modal_name": "添加全服邮件",
        "mail_name": this.getAllMailName(),
        "server_name": this.getServerName(),
        "onSend": this.OnSend
      }), React.createElement("hr", null), React.createElement("div", null, "全服邮件:"), React.createElement(MailList, Object.assign({}, this.props, {
        "ref": "mail_list_all",
        "server_name": this.getServerName(),
        "is_personal": false,
        "mail_name": this.getAllMailName()
      })), React.createElement("hr", null));
    }
  });

  MailOne = React.createClass({displayName: "MailOne",
    getInitialState: function() {
      return {
        select_server: this.props.curr_server,
        player_to_send: "0:0:1331"
      };
    },
    handleServerChange: function(data) {
      console.log(data);
      this.setState({
        select_server: data
      });
      return setTimeout(this.OnSend, 100);
    },
    handleUserChange: function(data) {
      console.log(data);
      if (data === "") {
        data = "请输入玩家Id";
      }
      return this.setState({
        player_to_send: data
      });
    },
    OnSend: function() {
      console.log("SendOver");
      return this.refs.mail_list.Refersh(1);
    },
    getAllMailName: function() {
      var res;
      if (this.state.select_server != null) {
        res = this.state.select_server.serverName.split(':');
        return "all:" + res[1];
      } else {
        return "all:10";
      }
    },
    getServerName: function() {
      if (this.state.select_server != null) {
        return this.state.select_server.serverName;
      } else {
        return "";
      }
    },
    render: function() {
      return React.createElement("div", null, React.createElement(AccountKeyInput, {
        "can_cb": this.handleUserChange
      }), React.createElement("div", {
        "className": "row-flex row-flex-middle row-flex-start"
      }, React.createElement(NewMailAllSendModal, {
        "modal_name": "添加单人邮件给:" + this.state.player_to_send,
        "mail_name": "profile:" + this.state.player_to_send,
        "server_name": this.getServerName(),
        "onSend": this.OnSend
      }), React.createElement("button", {
        "className": "ant-btn ant-btn-primary",
        "onClick": this.OnSend
      }, "查询个人所有邮件")), React.createElement("hr", null), React.createElement("div", null, "单人邮件[" + this.state.player_to_send + "]:"), React.createElement(MailList, Object.assign({
        "ref": "mail_list"
      }, this.props, {
        "server_name": this.getServerName(),
        "is_personal": true,
        "mail_name": "profile:" + this.state.player_to_send
      })), React.createElement("hr", null));
    }
  });

  MailBatch = React.createClass({displayName: "MailBatch",
    getInitialState: function() {
      var api;
      api = new Api();
      api.Typ("getBatchMailIndex").ServerID("").AccountID("").Key(this.props.curr_key).ParamArray([]).Do((function(_this) {
        return function(result) {
          return _this.setState({
            batch_index: "当前批量号: " + result
          });
        };
      })(this));
      return {
        batch_query_index: "",
        batch_index: ""
      };
    },
    onSuccess: function(data) {
      if (data.file.status === "done") {
        this.refs.batch_result.Refresh(data.file.response);
        this.setState({
          batch_index: "当前批量号: " + data.file.response.CurrentBatchIndex
        });
      }
      if (data.file.status === "error") {
        return alert(data.file.response);
      }
    },
    post: function(url, params) {
      var opt, temp, x;
      temp = document.createElement('form');
      temp.action = url;
      temp.method = 'post';
      temp.style.display = 'none';
      for (x in params) {
        opt = document.createElement('input');
        opt.name = x;
        opt.value = params[x];
        temp.appendChild(opt);
      }
      document.body.appendChild(temp);
      temp.submit();
      return temp;
    },
    query: function() {
      var data;
      data = {
        "params": this.params,
        "key": this.key
      };
      return this.post("../api/v1/command/null/null" + "/" + "getBatchMailDetail", {
        key: this.props.curr_key,
        params: [this.state.batch_query_index]
      });
    },
    onChange: function(e) {
      return this.setState({
        batch_query_index: e.target.value
      });
    },
    render: function() {
      return React.createElement("div", null, React.createElement("div", null, React.createElement(Input, {
        "value": this.state.batch_index,
        "disabled": "disabled"
      })), React.createElement("hr", null), React.createElement("div", null, React.createElement(Upload, Object.assign({}, this.props, {
        "action": "/api/v1/batchmail",
        "onChange": this.onSuccess
      }), React.createElement(Button, {
        "type": "ghost"
      }, React.createElement(Icon, {
        "type": "upload"
      }), " 批量发送邮件"))), React.createElement("hr", null), React.createElement("div", null, React.createElement(Upload, Object.assign({}, this.props, {
        "action": "/api/v1/batch_del_mail",
        "onChange": this.onSuccess
      }), React.createElement(Button, {
        "type": "ghost"
      }, React.createElement(Icon, {
        "type": "upload"
      }), " 批量删除邮件"))), React.createElement("hr", null), React.createElement("div", null, React.createElement(Input, {
        "onChange": this.onChange
      }), React.createElement(Button, {
        "bsStyle": 'primary',
        "onClick": this.query
      }, "下载批量邮件详情")), React.createElement("hr", null), React.createElement("div", null, React.createElement(BatchMailResult, {
        "ref": "batch_result"
      })));
    }
  });

  module.exports = {
    "MailAll": MailAll,
    "MailOne": MailOne,
    "MailBatch": MailBatch
  };

}).call(this);