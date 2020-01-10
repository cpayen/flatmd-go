const { src, dest, parallel, watch } = require('gulp');
var sass = require('gulp-sass');
var livereload = require('gulp-livereload');
var sourcemaps = require('gulp-sourcemaps');
var cleanCss = require('gulp-clean-css');
var ts = require('gulp-typescript');
var tsProject = ts.createProject('tsconfig.json');

var paths = {
    cssSrc: './scss/*.scss',
    cssDist: '../css',
    tsSrc: './ts/*.ts',
    tsDist: '../js'
}

function sassBuild() {
    return src(paths.cssSrc)
        .pipe(sass())
        .pipe(cleanCss())
        .pipe(dest(paths.cssDist));
}
exports.sassBuild = sassBuild;

function sassDev() {
    return src(paths.cssSrc)
        .pipe(sourcemaps.init())
        .pipe(sass().on('error', sass.logError))
        .pipe(sourcemaps.write())
        .pipe(dest(paths.cssDist))
        //.pipe(livereload());
}
exports.sassDev = sassDev;

function tsDev () {
    return src(paths.tsSrc)
        .pipe(sourcemaps.init())
        .pipe(tsProject())
        .js
        .pipe(sourcemaps.write())
        .pipe(dest(paths.tsDist));
};
exports.tsDev = tsDev;

function tsBuild () {
    return src(paths.tsSrc)
        .pipe(tsProject())
        .js
        .pipe(dest(paths.tsDist));
};
exports.tsBuild = tsBuild;

function watchDev () {
    livereload.listen(); 
    watch(['*.css', '*.html', '*.js']).on('change', livereload.changed);
    watch([paths.cssSrc], sassDev);
    watch([paths.tsSrc], tsDev);
}
exports.watchDev = watchDev;

exports.default = parallel(sassBuild, tsBuild);
exports.build = parallel(sassBuild, tsBuild);
exports.dev = parallel(sassDev, tsDev);



// gulp.task('watch', ['sass-dev'], function() {
// 	livereload.listen();
//     gulp.watch(['*.php', '*.html', '*.js']).on('change', livereload.changed);
// 	gulp.watch('styles/**', ['sass-dev']);
//     gulp.watch('scripts/*.ts', ['typescript-dev']);
// });