define(['exports', 'module', 'react', 'classnames', './BootstrapMixin'], function (exports, module, _react, _classnames, _BootstrapMixin) {
  'use strict';

  var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

  function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

  function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

  var _React = _interopRequireDefault(_react);

  var _classNames = _interopRequireDefault(_classnames);

  var _BootstrapMixin2 = _interopRequireDefault(_BootstrapMixin);

  var NavItem = _React['default'].createClass({
    displayName: 'NavItem',

    mixins: [_BootstrapMixin2['default']],

    propTypes: {
      onSelect: _React['default'].PropTypes.func,
      active: _React['default'].PropTypes.bool,
      disabled: _React['default'].PropTypes.bool,
      href: _React['default'].PropTypes.string,
      title: _React['default'].PropTypes.node,
      eventKey: _React['default'].PropTypes.any,
      target: _React['default'].PropTypes.string
    },

    getDefaultProps: function getDefaultProps() {
      return {
        href: '#'
      };
    },

    render: function render() {
      var _props = this.props;
      var disabled = _props.disabled;
      var active = _props.active;
      var href = _props.href;
      var title = _props.title;
      var target = _props.target;
      var children = _props.children;

      var props = _objectWithoutProperties(_props, ['disabled', 'active', 'href', 'title', 'target', 'children']);

      // eslint-disable-line object-shorthand
      var classes = {
        active: active,
        disabled: disabled
      };
      var linkProps = {
        href: href,
        title: title,
        target: target,
        onClick: this.handleClick,
        ref: 'anchor'
      };

      if (href === '#') {
        linkProps.role = 'button';
      }

      return _React['default'].createElement(
        'li',
        _extends({}, props, { className: (0, _classNames['default'])(props.className, classes) }),
        _React['default'].createElement(
          'a',
          linkProps,
          children
        )
      );
    },

    handleClick: function handleClick(e) {
      if (this.props.onSelect) {
        e.preventDefault();

        if (!this.props.disabled) {
          this.props.onSelect(this.props.eventKey, this.props.href, this.props.target);
        }
      }
    }
  });

  module.exports = NavItem;
});