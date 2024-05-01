fitcha
======

Small feature management library for Go developers. Slightly opionated, highly experimental. YMMV.

Fitcha enables you to check features available to users per-request or globally. It stores 
features in a `Store`, and uses `Manager` to check availability of a feature. Features can
be added to the store dynamically or pre-loaded. One of the things that makes fitcha special
is it's support for dynamic expressions which enable you to implement feature validation with
additional logic via expressions written with [expr](https://github.com/expr-lang/expr) - hopefully this will give you more power for rolling out features.

The API is unstable, the project is experimental - but we are dogfooding it so if doesn't work
out, we have a lot of refactoring to do on our internal projects :D
