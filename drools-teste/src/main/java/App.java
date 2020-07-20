
import org.kie.api.KieServices;
import org.kie.api.runtime.KieContainer;
import org.kie.api.runtime.KieSession;

import models.Person;

public class App {
    public static void main(String[] args) {
        System.out.println("Iniciando a engine de regras do Drools");
        KieServices ks = KieServices.Factory.get();
        KieContainer kieContainer = ks.getKieClasspathContainer();

        KieSession kSession = kieContainer.newKieSession();

        Person p = new Person("");

        kSession.insert(p);

        int fired = kSession.fireAllRules();

        System.out.println("Numero de regras executadas: " + fired);
        System.out.println("Nome da pessoa: " + p.getName());
    }
}
