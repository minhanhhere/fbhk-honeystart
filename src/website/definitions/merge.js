F.config['allow-compile-script'] = false;

F.merge('/js/default.js', '/js/jctajr.min.js', '/js/ui.js', '/js/default.js');

var fs = require('fs'),
    path = require('path'),
    less = require('less');

var src = __dirname + "/../public/styles/app-base.less";
var dest = __dirname + "/../public/styles/main.css";
less.render(fs.readFileSync(src).toString(), {
    filename: path.resolve(src),
    compress: true,
}, function (e, output) {
    fs.writeFile(dest, output.css, function (err) {
        if (err) {
            console.log(err);
        }
    });
});

F.merge('/styles/vendor.css',
    '../bower_components/angular-loading-bar/build/loading-bar.css',
    '../bower_components/animate.css/animate.css',
    '../bower_components/hover/css/hover.css',
    '../bower_components/textAngular/dist/textAngular.css',
    '../bower_components/AngularJS-Toaster/toaster.css',
    '../bower_components/font-awesome/css/font-awesome.min.css',
    '../bower_components/ladda/dist/ladda-themeless.min.css',
    '/styles/loaders.min.css'
);

F.merge('/scripts/vendor.js',
    '/scripts/jctajr.min.js',
    '../bower_components/jquery/dist/jquery.min.js',
    '/scripts/Extras/modernizr.custom.js',
    '../bower_components/angular/angular.js',
    '../bower_components/angular-animate/angular-animate.js',
    '../bower_components/angular-cookies/angular-cookies.js',
    '../bower_components/angular-growl/build/angular-growl.js',
    '../bower_components/angular-growl-notifications/dist/angular-growl-notifications.js',
    '../bower_components/angular-loading-bar/build/loading-bar.js',
    '../bower_components/angular-mocks/angular-mocks.js',
    '../bower_components/angular-progress-button-styles/dist/angular-progress-button-styles.min.js',
    '../bower_components/angular-resource/angular-resource.js',
    '../bower_components/angular-translate/angular-translate.js',
    '../bower_components/angular-translate-loader-static-files/angular-translate-loader-static-files.js',
    '../bower_components/angular-translate-loader-url/angular-translate-loader-url.js',
    '../bower_components/moment/moment.js',
    '../bower_components/angular-ui-router/release/angular-ui-router.js',
    '../bower_components/classie/classie.js',
    '../bower_components/moment-timezone/builds/moment-timezone-with-data-2010-2020.js',
    '../bower_components/lodash/lodash.js',
    '../bower_components/rangy/rangy-core.js',
    '../bower_components/rangy/rangy-classapplier.js',
    '../bower_components/rangy/rangy-highlighter.js',
    '../bower_components/rangy/rangy-selectionsaverestore.js',
    '../bower_components/rangy/rangy-serializer.js',
    '../bower_components/rangy/rangy-textrange.js',
    '../bower_components/textAngular/dist/textAngular.js',
    '../bower_components/textAngular/dist/textAngular-sanitize.js',
    '../bower_components/textAngular/dist/textAngularSetup.js',
    '../bower_components/angular-breadcrumb/release/angular-breadcrumb.js',
    '../bower_components/AngularJS-Toaster/toaster.js',
    '../bower_components/angular-sanitize/angular-sanitize.js',
    '../bower_components/humanize-duration/humanize-duration.js',
    '../bower_components/angular-timer/dist/angular-timer.js',
    '../bower_components/perfect-scrollbar/js/perfect-scrollbar.jquery.js',
    '../bower_components/angular-bootstrap/ui-bootstrap-tpls.min.js',
    '../bower_components/moment/min/moment.min.js',
    '../bower_components/eonasdan-bootstrap-datetimepicker/build/js/bootstrap-datetimepicker.min.js',
    '../bower_components/spin.js/spin.js',
    '../bower_components/ladda/dist/ladda.min.js',
    '../bower_components/ladda-angular/dist/ladda-angular.min.js',
    '../bower_components/wow/dist/wow.min.js',
    '../bower_components/angular-avatar/dist/angular-avatar.js'
);

F.merge('/scripts/oldieshim.js',
    '../bower_components/es5-shim/es5-shim.js',
    '../bower_components/json3/lib/json3.js'
);

F.merge('/scripts/main.js',
    '/scripts/Business/rest-module.js',
    '/scripts/Home/home-ctrl.js',
    '/scripts/Project/project-ctrl.js',
    '/scripts/Project/project-new-ctrl.js',
    '/scripts/Project/picker-modal.js',
    '/scripts/Payment/payment-ctrl.js',
    '/scripts/Payment/success-ctrl.js',
    '/scripts/global.js',
    '/scripts/app.js',
    '/scripts/Auth/auth-modal.js',
    '/scripts/Auth/auth-service.js',
    '/scripts/Extras/progress-button.js',
    '/scripts/Autofocus/autofocus.js',
    '/scripts/TopNav/topnav-ctrl.js',
    '/scripts/Comment/disqus.js'
);