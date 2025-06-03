package com.mycompany

class Utils implements Serializable {

    def steps

    Utils(steps) {
        this.steps = steps
    }

    void printInfo(String msg) {
        steps.echo "[INFO] ${msg}"
    }

    boolean isProductionEnv(String env) {
        return env?.toLowerCase() == "production"
    }
}
