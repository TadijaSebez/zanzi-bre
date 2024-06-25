# OWASP ASVS 4.0.3 analiza projekta

## V1 Architecture, Design and Threat Modeling

### V1.1 Secure Software Development Lifecycle

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.1.1|Verify the use of a secure software development lifecycle that addresses security in all stages of development. (C1)||&check;|&check;||Da|Radili smo security code review, imali smo CI pipeline|
|1.1.2|Verify the use of threat modeling for every design change or sprint planningto identify threats, plan for countermeasures, facilitate appropriate risk responses, and guide security testing.||&check;|&check;|1053|Da|Napravili smo threat model|
|1.1.3|Verify that all user stories and features contain functional security constraints, such as "As a user, I should be able to view and edit my profile. I should not be able to view or edit anyone else's profile"||&check;|&check;|1110|Ne|Nismo koristili user storyje|
|1.1.4|Verify documentation and justification of all the application's trust boundaries, components, and significant data flows.||&check;|&check;|1059|Ne||
|1.1.5|Verify definition and security analysis of the application's high-level architecture and all connected remote services. (C1)||&check;|&check;|1059|Da|Uradili smo analizu bezbednosti|
|1.1.6|Verify implementation of centralized, simple (economy of design), vetted, secure, and reusable security controls to avoid duplicate, missing, ineffective, or insecure controls. (C10)||&check;|&check;|637|Da||
|1.1.7|Verify availability of a secure coding checklist, security requirements, guideline, or policy to all developers and testers.||&check;|&check;|637|Ne|Nismo napravili checklistu za bezbednost koda|

Projekat je jednostavan pa nema potrebe za ispunjenjem stavki koje nisu ispunjene. Poglavlje V1.1 je zadovoljeno u dovoljnoj meri.


### V1.2 Authentication Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.2.1|Verify the use of unique or special low-privilege operating system accounts for all application components, services, and servers. (C3)||&check;|&check;|250|Da|Svaka komponenta radi bez administratorskih privilegija|
|1.2.2|Verify that communications between application components, including APIs, middleware and data layers, are authenticated. Components should have the least necessary privileges needed. (C3)||&check;|&check;|306|Ne|Nema autentifikacije izmedju servisnog sloja i sloja podataka. Komponente imaju podrazumevane privilegije|
|1.2.3|Verify that the application uses a single vetted authentication mechanism that is known to be secure, can be extended to include strong authentication, and has sufficient logging and monitoring to detect account abuse or breaches||&check;|&check;|306|Ne|Ne logujemo dovoljno poruka i nemamo mehanizam monitoringa logova|
|1.2.4|Verify that all authentication pathways and identity management APIs implement consistent authentication security control strength, such that there are no weaker alternatives per the risk of the application.||&check;|&check;|306|Da|Nemamo alternativnih puteva za autentifikaciju koji su slabiji|

Trebalo bi poboljšati sigurnost komunikacije između komponenti, poboljšati logovanje i dodati monitoring logova.


### V1.3 Session Management Architecture

Ovo poglavlje je samo placeholder za novije verzije OWASP ASVS.


### V1.4 Access Control Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.4.1|Verify that trusted enforcement points, such as access control gateways, servers, and serverless functions, enforce access controls. Never enforce access controls on the client.||&check;|&check;|602|Da||
|1.4.2|[DELETED, NOT ACTIONABLE]|||||||
|1.4.3|[DELETED, DUPLICATE OF 4.1.3]||||||
|1.4.4|Verify the application uses a single and well-vetted access control mechanism for accessing protected data and resources. All requests must pass through this single mechanism to avoid copy and paste or insecure alternative paths. (C7)||&check;|&check;|284|Da|Minimalistička verzija google zanzibara kontroliše pristup podacima i resursima|
|1.4.5|Verify that attribute or feature-based access control is used whereby the code checks the user's authorization for a feature/data item rather than just their role. Permissions should still be allocated using roles. (C7)||&check;|&check;|275|Da||

Ovo poglavlje je zadovoljneno.


### V1.5 Input and Output Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.5.1|Verify that input and output requirements clearly define how to handle and process data based on type, content, and applicable laws, regulations, and other policy compliance.||&check;|&check;|1029|Ne|Nismo ovo razmatrali|
|1.5.2|Verify that serialization is not used when communicating with untrusted clients. If this is not possible, ensure that adequate integrity controls (and possibly encryption if sensitive data is sent) are enforced to prevent deserialization attacks including object injection.||&check;|&check;|502|Da|Nije moguće uraditi deserialization attack|
|1.5.3|Verify that input validation is enforced on a trusted service layer. (C5)||&check;|&check;|602|Da|Podaci se validiraju na serveru|
|1.5.4|Verify that output encoding occurs close to or by the interpreter for which it is intended. (C4)||&check;|&check;|116|Da||

Trebalo bi da razmotrimo zakone i propise koji se odnose na naše čuvanje podataka.


### V1.6 Cryptographic Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.6.1|Verify that there is an explicit policy for management of cryptographic keys and that a cryptographic key lifecycle follows a key management standard such as NIST SP 800-57.||&check;|&check;|320|Ne|Nismo osmislili pravila za upravljanjem kriptografskim ključevima|
|1.6.2|Verify that consumers of cryptographic services protect key material and other secrets by using key vaults or API based alternatives.||&check;|&check;|320|Ne||
|1.6.3|Verify that all keys and passwords are replaceable and are part of a welldefined process to re-encrypt sensitive data.||&check;|&check;|320|Ne|Ne podržavamo menjanje šifri. Na primer šifra koja se koristi za generisanje JWT tokena je hardcodovana|
|1.6.4|Verify that the architecture treats client-side secrets--such as symmetric keys, passwords, or API tokens--as insecure and never uses them to protect or access sensitive data.||&check;|&check;|320|Da|Client-side secret se koristi samo za autentifikaciju|

Trebalo bi osmisliti mehanizam za upravljanje šiframa i kriptografskim ključevima.


### V1.7 Errors, Logging and Auditing Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.7.1|Verify that a common logging format and approach is used across the system. (C9)||&check;|&check;|1009|Da|Isti format i pristup su korišćeni u celoj aplikaciji|
|1.7.2|Verify that logs are securely transmitted to a preferably remote system for analysis, detection, alerting, and escalation. (C9)||&check;|&check;||Ne|Logovi se ne šalju na monitoring i analizu|

Trebalo bi implementirati monitoring logova, bolje mehanizme čuvanja i upravljanja logovima.


### V1.8 Data Protection and Privacy Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.8.1|Verify that all sensitive data is identified and classified into protection levels.||&check;|&check;||Ne|Nismo definisali nivoe zaštite za osetljive podatke|
|1.8.2|Verify that all protection levels have an associated set of protection requirements, such as encryption requirements, integrity requirements, retention, privacy and other confidentiality requirements, and that these are applied in the architecture.||&check;|&check;||Ne||


### V1.9 Communications Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.9.1|Verify the application encrypts communications between components, particularly when these components are in different containers, systems, sites, or cloud providers. (C3)||&check;|&check;|319|Ne|Komunikacija sa bazom podataka nije enkriptovana, ali komunikacija izmađu ostalih komponenti jeste|
|1.9.2|Verify that application components verify the authenticity of each side in a communication link to prevent person-in-the-middle attacks. For example, application components should validate TLS certificates and chains||&check;|&check;|295|Da|Sertifikati se validiraju kod https komunikacije|

Trebalo bi enkriptovati komunikaciju sa bazom podataka.


### V1.10 Malicious Software Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.10.1|Verify that a source code control system is in use, with procedures to ensure that check-ins are accompanied by issues or change tickets. The source code control system should have access control and identifiable users to allow traceability of any changes.||&check;|&check;|284|Da|Korišćeni su git i github, ali nisu korišćeni issues|


### V1.11 Business Logic Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.11.1|Verify the definition and documentation of all application components interms of the business or security functions they provide.||&check;|&check;|1059|Da||
|1.11.2|Verify that all high-value business logic flows, including authentication, session management and access control, do not share unsynchronized state.||&check;|&check;|362|Da||
|1.11.3|Verify that all high-value business logic flows, including authentication, session management and access control are thread safe and resistant to time-of-check and time-of-use race conditions.|||&check;|367|Da||


### V1.12 Secure File Upload Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.12.1|[DELETED, DUPLICATE OF 12.4.1]||||||
|1.12.2|Verify that user-uploaded files - if required to be displayed or downloaded from the application - are served by either octet stream downloads, or from an unrelated domain, such as a cloud file storage bucket. Implement a suitable Content Security Policy (CSP) to reduce the risk from XSS vectors or other attacks from the uploaded file.||&check;|&check;|Ne|Nemamo cloud file storage bucket|


### V1.13 API Architecture

Ovo poglavlje je samo placeholder za novije verzije OWASP ASVS.


### V1.14 Configuration Architecture

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|1.14.1|Verify the segregation of components of differing trust levels through welldefined security controls, firewall rules, API gateways, reverse proxies, cloud-based security groups, or similar mechanisms.||&check;|&check;|923|Ne||
|1.14.2|Verify that binary signatures, trusted connections, and verified endpoints are used to deploy binaries to remote devices.||&check;|&check;|494|Da|Ne radimo deploy binarnih fajlova|
|1.14.3|Verify that the build pipeline warns of out-of-date or insecure components and takes appropriate actions.||&check;|&check;|1104|Da|Imamo relevantne provere u CI pipeline|
|1.14.4|Verify that the build pipeline contains a build step to automatically build and verify the secure deployment of the application, particularly if the application infrastructure is software defined, such as cloud environment build scripts.||&check;|&check;||Ne|Nismo deployovali aplikaciju tako da to još uvek ne postoji u build pipelineu|
|1.14.5|Verify that application deployments adequately sandbox, containerize and/or isolate at the network level to delay and deter attackers from attacking other applications, especially when they are performing sensitive or dangerous actions such as deserialization. (C5)||&check;|&check;|265|Ne||
|1.14.6|Verify the application does not use unsupported, insecure, or deprecated client-side technologies such as NSAPI plugins, Flash, Shockwave, ActiveX, Silverlight, NACL, or client-side Java applets.||&check;|&check;|477|Da|Ne koristimo zastarele nebezbedne tehnologije|

Trebalo bi uraditi bezbedan deploy aplikacije.

