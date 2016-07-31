'use strict';

angular
    .module('HoneyStart', [
        'ngAnimate',
        //'ui.calendar',
        //'chart.js',
        //'gridshore.c3js.chart',
        'angular-growl',
        'growlNotifications',
        'angular-loading-bar',
        'angular-progress-button-styles',
        //'pascalprecht.translate',
        'ui.bootstrap',
        'ngCookies',
        //'ngResource',
        'ngSanitize',
        'timer',
        'ncy-angular-breadcrumb',
        'toaster',
        'ladda',
        'app.home',
        'app.project',
        'app.payment',
        'app.modal',
        'app.auth',
        'ngAvatar'
    ])
    .config(['cfpLoadingBarProvider', function (cfpLoadingBarProvider) {
        cfpLoadingBarProvider.latencyThreshold = 100;
        cfpLoadingBarProvider.includeSpinner = true;
    }])
    //.config(['$translateProvider', function ($translateProvider) {
    //    $translateProvider.useStaticFilesLoader({
    //        prefix: '/languages/',
    //        suffix: '.json'
    //    });
    //    $translateProvider.preferredLanguage('vn');
    //}])
;
