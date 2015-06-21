(ns musician.core
  (:use [overtone.live])
  (:gen-class))

(defn -main [& args]
  (println "Testing overtone")
  (println "This should last for about 10 seconds.")

  ;; start up some noise
  (demo 10 (* (sin-osc 0.5) (saw 440)))

  ;; wait around...
  (Thread/sleep (* 10 1000))

  (println "Done. The program should stop.")
  (System/exit 0))
