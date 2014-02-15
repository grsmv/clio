module.exports = function (grunt) {
    grunt.initConfig({
        pkg: grunt.file.readJSON("package.json"),

        sass: {
            dist: {
                options: {
                    style: "expanded"
                },
                files: {
                    "css/styles.css": "scss/styles.scss"
                }
            }
        },

        cssmin: {
            combine: {
                files: {
                    "css/styles.min.css": [
                        "lib/bootstrap/bootstrap.css",
                        "lib/normalize-css/normalize.css",
                        "css/styles.css"
                    ]
                }
            }
        },

        bower: {
            install: {
                options: {
                    targetDir: "./lib",
                    layout: "byType",
                    install: true,
                    verbose: false,
                    cleanTargetDir: false,
                    cleanBowerDir: false,
                    bowerOptions: {}
                }
            }
        },

        "closure-compiler": {
            frontend: {
                closurePath: "$HOME/bin/closure-compiler",
                js: [
                    "lib/jquery/jquery.js",
                    "lib/bootstrap/bootstrap.js"
                ],
                jsOutputFile: "javascripts/all.min.js",
                maxBuffer: 500,
                options: {
                    compilation_level: "SIMPLE_OPTIMIZATIONS"
                }
            }
        },

        watch: {
            options: {
                livereload: false
            },
            scss: {
                files: [
                    "scss/styles.scss",
                ],
                tasks: ["sass", "cssmin"]
            },
        }
    });

    grunt.loadNpmTasks("grunt-bower-task");
    grunt.loadNpmTasks("grunt-closure-compiler");
    grunt.loadNpmTasks("grunt-contrib-cssmin");
    grunt.loadNpmTasks("grunt-contrib-sass");
    grunt.loadNpmTasks("grunt-contrib-watch");

    grunt.registerTask("watch_scss", ["watch:scss"]);
};
