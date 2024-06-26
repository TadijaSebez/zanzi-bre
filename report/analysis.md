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


## V2 Authentication

### V2.1 Password Security

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.1.1|Verify that user set passwords are at least 12 characters in length (after multiple spaces are combined). (C6)|&check;|&check;|&check;|521|Ne|Nema provere za dužinu passworda na backu|
|2.1.2|Verify that passwords of at least 64 characters are permitted, and that passwords of more than 128 characters are denied. (C6)|&check;|&check;|&check;|521|Ne||
|2.1.3|Verify that password truncation is not performed. However, consecutive multiple spaces may be replaced by a single space. (C6)|&check;|&check;|&check;|521|Da||
|2.1.4|Verify that any printable Unicode character, including language neutral characters such as spaces and Emojis are permitted in passwords.|&check;|&check;|&check;|521|Da||
|2.1.5|Verify users can change their password.|&check;|&check;|&check;|620|Ne|Nije implementirano|
|2.1.6|Verify that password change functionality requires the user's current and new password.|&check;|&check;|&check;|620|Ne||
|2.1.7|Verify that passwords submitted during account registration, login, and password change are checked against a set of breached passwords either locally (such as the top 1,000 or 10,000 most common passwords which match the system's password policy) or using an external API. If using an API a zero knowledge proof or other mechanism should be used to ensure that the plain text password is not sent or used in verifying the breach status of the password. If the password is breached, the application must require the user to set a new non-breached password. (C6)|&check;|&check;|&check;|521|Ne|Nije implementirano|
|2.1.8|Verify that a password strength meter is provided to help users set a stronger password.|&check;|&check;|&check;|521|Ne||
|2.1.9|Verify that there are no password composition rules limiting the type of characters permitted. There should be no requirement for upper or lower case or numbers or special characters. (C6)|&check;|&check;|&check;|521|Da|Nema pravila koja ograničavaju karaktere u passwordu|
|2.1.10|Verify that there are no periodic credential rotation or password history requirements.|&check;|&check;|&check;|263|Da|Nema rotacije passworda|
|2.1.11|Verify that "paste" functionality, browser password helpers, and external password managers are permitted.|&check;|&check;|&check;|521|Da||
|2.1.12|Verify that the user can choose to either temporarily view the entire masked password, or temporarily view the last typed character of the password on platforms that do not have this as built-in functionality.|&check;|&check;|&check;|521|Da||

Trebalo bi ograničiti dužinu passworda, zahtevati jake passworde koji nisu ukradeni u prethodnim hakerskim napadima i dozvoliti promenu passworda.


### V2.2 General Authenticator Security

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.2.1|Verify that anti-automation controls are effective at mitigating breached credential testing, brute force, and account lockout attacks. Such controls include blocking the most common breached passwords, soft lockouts, rate limiting, CAPTCHA, ever increasing delays between attempts, IP address restrictions, or risk-based restrictions such as location, first login on a device, recent attempts to unlock the account, or similar. Verify that no more than 100 failed attempts per hour is possible on a single account.|&check;|&check;|&check;|307|Ne||
|2.2.2|Verify that the use of weak authenticators (such as SMS and email) is limited to secondary verification and transaction approval and not as a replacement for more secure authentication methods. Verify that stronger methods are offered before weak methods, users are aware of the risks, or that proper measures are in place to limit the risks of account compromise.|&check;|&check;|&check;|304|Da|Nemamo SMS ili email autentifikaciju|
|2.2.3|Verify that secure notifications are sent to users after updates toauthentication details, such as credential resets, email or address changes, logging in from unknown or risky locations. The use of push notifications - rather than SMS or email - is preferred, but in the absence of push notifications, SMS or email is acceptable as long as no sensitive information is disclosed in the notification.|&check;|&check;|&check;|620|Ne|Nemamo notifikacije|
|2.2.4|Verify impersonation resistance against phishing, such as the use of multi-factor authentication, cryptographic devices with intent (such as connected keys with a push to authenticate), or at higher AAL levels, client-side certificates.|||&check;|308|Ne|Nemamo multi-factor autentifikaciju|
|2.2.5|Verify that where a Credential Service Provider (CSP) and the application verifying authentication are separated, mutually authenticated TLS is in place between the two endpoints.|||&check;|319|Da|Imamo https komunikaciju između komponenti|
|2.2.6|Verify replay resistance through the mandated use of One-time Passwords (OTP) devices, cryptographic authenticators, or lookup codes.|||&check;|308|Ne|Nemamo one-time passworde|
|2.2.7|Verify intent to authenticate by requiring the entry of an OTP token or user-initiated action such as a button press on a FIDO hardware key.|||&check;|308|Ne||

Trebalo bi dodati notifikacije za promene passworda, dodati multi-factor autentifikaciju i one-time passworde.


### V2.3 Authenticator Lifecycle

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.3.1|Verify system generated initial passwords or activation codes SHOULD be securely randomly generated, SHOULD be at least 6 characters long, and MAY contain letters and numbers, and expire after a short period of time. These initial secrets must not be permitted to become the long term password.|&check;|&check;|&check;|330|Da|Nemamo inicijalne passworde generisane od sistema|
|2.3.2|Verify that enrollment and use of user-provided authentication devices are supported, such as a U2F or FIDO tokens.||&check;|&check;|308|Ne||
|2.3.3|Verify that renewal instructions are sent with sufficient time to renew time bound authenticators.||&check;|&check;|287|Ne||


### V2.4 Credential Storage

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.4.1|Verify that passwords are stored in a form that is resistant to offline attacks. Passwords SHALL be salted and hashed using an approved one-way key derivation or password hashing function. Key derivation and password hashing functions take a password, a salt, and a cost factor as inputs when generating a password hash. (C6)||&check;|&check;|916|Ne|Nismo koristili salt|
|2.4.2|Verify that the salt is at least 32 bits in length and be chosen arbitrarily to minimize salt value collisions among stored hashes. For each credential, a unique salt value and the resulting hash SHALL be stored. (C6)||&check;|&check;|916|Ne||
|2.4.3|Verify that if PBKDF2 is used, the iteration count SHOULD be as large as verification server performance will allow, typically at least 100,000 iterations. (C6)||&check;|&check;|916|Da|Nismo koristili PBKDF2|
|2.4.4|Verify that if bcrypt is used, the work factor SHOULD be as large as verification server performance will allow, with a minimum of 10. (C6)||&check;|&check;|916|Da||
|2.4.5|Verify that an additional iteration of a key derivation function is performed, using a salt value that is secret and known only to the verifier. Generate the salt value using an approved random bit generator [SP 800-90Ar1] and provide at least the minimum security strength specified in the latest revision of SP 800-131A. The secret salt value SHALL be stored separately from the hashed passwords (e.g., in a specialized device like a hardware security module).||&check;|&check;|916|Ne||

Trebalo bi dodati salt kod heshiranja passworda.


### V2.5 Credential Recovery

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.5.1|Verify that a system generated initial activation or recovery secret is not sent in clear text to the user. (C6)|&check;|&check;|&check;|640|Da||
|2.5.2|Verify password hints or knowledge-based authentication (socalled "secret questions") are not present.|&check;|&check;|&check;|640|Da||
|2.5.3|Verify password credential recovery does not reveal the current password in any way. (C6)|&check;|&check;|&check;|640|Da||
|2.5.4|Verify shared or default accounts are not present (e.g. "root", "admin", or "sa").|&check;|&check;|&check;|16|Da||
|2.5.5|Verify that if an authentication factor is changed or replaced, that the user is notified of this event.|&check;|&check;|&check;|304|Ne|Nemamo notifikacije|
|2.5.6|Verify forgotten password, and other recovery paths use a secure recovery mechanism, such as time-based OTP (TOTP) or other soft token, mobile push, or another offline recovery mechanism. (C6)|&check;|&check;|&check;|640|Ne||
|2.5.7|Verify that if OTP or multi-factor authentication factors are lost, that evidence of identity proofing is performed at the same level as during enrollment.|&check;|&check;|&check;|308|Ne||


### V2.6 Look-up Secret Verifier

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.6.1|Verify that lookup secrets can be used only once.||&check;|&check;|308|Da|Nemamo lookup secret|
|2.6.2|Verify that lookup secrets have sufficient randomness (112 bits of entropy), or if less than 112 bits of entropy, salted with a unique and random 32-bit salt and hashed with an approved one-way hash.||&check;|&check;|330|Da||
|2.6.3|Verify that lookup secrets are resistant to offline attacks, such as predictable values.||&check;|&check;|310|Da||


### V2.7 Out of Band Verifier

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.7.1|Verify that clear text out of band (NIST "restricted") authenticators, such as SMS or PSTN, are not offered by default, and stronger alternatives such as push notifications are offered first.|&check;|&check;|&check;|287|Ne|Nemamo push notifikacije|
|2.7.2|Verify that the out of band verifier expires out of band authentication requests, codes, or tokens after 10 minutes.|&check;|&check;|&check;|287|Ne||
|2.7.3|Verify that the out of band verifier authentication requests, codes, or tokens are only usable once, and only for the original authentication request.|&check;|&check;|&check;|287|Ne||
|2.7.4|Verify that the out of band authenticator and verifier communicates over a secure independent channel.|&check;|&check;|&check;|523|Ne||
|2.7.5|Verify that the out of band verifier retains only a hashed version of the authentication code.||&check;|&check;|256|Ne||
|2.7.6|Verify that the initial authentication code is generated by a secure random number generator, containing at least 20 bits of entropy (typically a six digital random number is sufficient).||&check;|&check;|310|Ne||


### V2.8 One Time Verifier

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.8.1|Verify that time-based OTPs have a defined lifetime before expiring.|&check;|&check;|&check;|613|Ne|Nemamo OTP|
|2.8.2|Verify that symmetric keys used to verify submitted OTPs are highly protected, such as by using a hardware security module or secure operating system based key storage.||&check;|&check;|320|Ne||
|2.8.3|Verify that approved cryptographic algorithms are used in the generation, seeding, and verification of OTPs.||&check;|&check;|326|Ne||
|2.8.4|Verify that time-based OTP can be used only once within the validity period.||&check;|&check;|287|Ne||
|2.8.5|Verify that if a time-based multi-factor OTP token is re-used during the validity period, it is logged and rejected with secure notifications being sent to the holder of the device.||&check;|&check;|287|Ne||
|2.8.6|Verify physical single-factor OTP generator can be revoked in case of theft or other loss. Ensure that revocation is immediately effective across logged in sessions, regardless of location.||&check;|&check;|613|Ne||
|2.8.7|Verify that biometric authenticators are limited to use only as secondary factors in conjunction with either something you have and something you know.|||&check;|308|Ne|Nemamo biometrijsku autentifikaciju|


### V2.9 Cryptographic Verifier

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.9.1|Verify that cryptographic keys used in verification are stored securely and protected against disclosure, such as using a Trusted Platform Module (TPM) or Hardware Security Module (HSM), or an OS service that can use this secure storage.||&check;|&check;|320|Ne||
|2.9.2|Verify that the challenge nonce is at least 64 bits in length, and statistically unique or unique over the lifetime of the cryptographic device.||&check;|&check;|330|Ne||
|2.9.3|Verify that approved cryptographic algorithms are used in the generation, seeding, and verification.||&check;|&check;|327|Ne||


### V2.10 Service Authentication

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|2.10.1|Verify that intra-service secrets do not rely on unchanging credentials such as passwords, API keys or shared accounts with privileged access.||&check;|&check;|287|Da||
|2.10.2|Verify that if passwords are required for service authentication, the service account used is not a default credential. (e.g. root/root or admin/admin are default in some services during installation).||&check;|&check;|255|Da||
|2.10.3|Verify that passwords are stored with sufficient protection to prevent offline recovery attacks, including local system access.||&check;|&check;|522|Ne||
|2.10.4|Verify passwords, integrations with databases and thirdparty systems, seeds and internal secrets, and API keys are managed securely and not included in the source code or stored within source code repositories. Such storage SHOULD resist offline attacks. The use of a secure software key store (L1), hardware TPM, or an HSM (L3) is recommended for password storage.||&check;|&check;|798|Ne|Passwordi se nalaze u repozitorijumu|

Trebalo bi više voditi računa o intraservisnim kredencijalima.


## V3 Session Management

### V3.1 Fundamental Session Management Security

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|3.1.1|Verify the application never reveals session tokens in URL parameters.|&check;|&check;|&check;|598|Da||


### V3.2 Session Binding

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|3.2.1|Verify the application generates a new session token on user authentication. (C6)|&check;|&check;|&check;|384|Da|Pravi se novi JWT token prilikom svake autentifikacije|
|3.2.2|Verify that session tokens possess at least 64 bits of entropy. (C6)|&check;|&check;|&check;|331|Da||
|3.2.3|Verify the application only stores session tokens in the browser using secure methods such as appropriately secured cookies (see section 3.4) or HTML 5 session storage.|&check;|&check;|&check;|539|Ne|Koristimo local storage|
|3.2.4|Verify that session tokens are generated using approved cryptographic algorithms. (C6)||&check;|&check;|331|Da||


### V3.3 Session Termination

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|3.3.1|Verify that logout and expiration invalidate the session token, such that the back button or a downstream relying party does not resume an authenticated session, including across relying parties. (C6)|&check;|&check;|&check;|613|Da||
|3.3.2|If authenticators permit users to remain logged in, verify that re-authentication occurs periodically both when actively used or after an idle period. (C6)|&check;|&check;|&check;|613|Ne|Ne koristimo refresh token|
|3.3.3|Verify that the application gives the option to terminate all other active sessions after a successful password change (including change via password reset/recovery), and that this is effective across the application, federated login (if present), and any relying parties.||&check;|&check;|613|Ne|Nemamo ovu opciju|
|3.3.4|Verify that users are able to view and (having re-entered login credentials) log out of any or all currently active sessions and devices.||&check;|&check;|613|Da||


### V3.4 Cookie-based Session Management

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|3.4.1|Verify that cookie-based session tokens have the 'Secure' attribute set. (C6)|&check;|&check;|&check;|614|Ne|Ne koristimo cookie|
|3.4.2|Verify that cookie-based session tokens have the 'HttpOnly' attribute set. (C6)|&check;|&check;|&check;|1004|Ne||
|3.4.3|Verify that cookie-based session tokens utilize the 'SameSite' attribute to limit exposure to cross-site request forgery attacks. (C6)|&check;|&check;|&check;|16|Ne||
|3.4.4|Verify that cookie-based session tokens use the "__Host-" prefix so cookies are only sent to the host that initially set the cookie.|&check;|&check;|&check;|16|Ne||
|3.4.5|Verify that if the application is published under a domain name with other applications that set or use session cookies that might disclose the session cookies, set the path attribute in cookie-based session tokens using the most precise path possible. (C6)|&check;|&check;|&check;|16|Ne||


### V3.5 Token-based Session Management

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|3.5.1|Verify the application allows users to revoke OAuth tokens that form trust relationships with linked applications.||&check;|&check;|290|Ne|Nemamo OAuth|
|3.5.2|Verify the application uses session tokens rather than static API secrets and keys, except with legacy implementations.||&check;|&check;|798|Da|Koristimo JWT|
|3.5.3|Verify that stateless session tokens use digital signatures, encryption, and other countermeasures to protect against tampering, enveloping, replay, null cipher, and key substitution attacks.||&check;|&check;|345|Da|Koristimo JWT|

Mogli bismo dodati i OAuth kao mogući način autentifikacije.


### V3.6 Federated Re-authentication

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|3.6.1|Verify that Relying Parties (RPs) specify the maximum authentication time to Credential Service Providers (CSPs) and that CSPs reauthenticate the user if they haven't used a session within that period.|||&check;|613|Ne||
|3.6.2|Verify that Credential Service Providers (CSPs) inform Relying Parties (RPs) of the last authentication event, to allow RPs to determine if they need to re-authenticate the user.|||&check;|613|Ne||


### V3.7 Defenses Against Session Management Exploits

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|3.7.1|Verify the application ensures a full, valid login session or requires reauthentication or secondary verification before allowing any sensitive transactions or account modifications.|&check;|&check;|&check;|306|Da|Nema senzitivnih transakcija|


## V4 Access Control

### V4.1 General Access Control Design

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|4.1.1|Verify that the application enforces access control rules on a trusted service layer, especially if client-side access control is present and could be bypassed.|&check;|&check;|&check;|602|Da|Access control je na servisnom sloju|
|4.1.2|Verify that all user and data attributes and policy information used by access controls cannot be manipulated by end users unless specifically authorized.|&check;|&check;|&check;|639|Da||
|4.1.3|Verify that the principle of least privilege exists - users should only be able to access functions, data files, URLs, controllers, services, and other resources, for which they possess specific authorization. This implies protection against spoofing and elevation of privilege. (C7)|&check;|&check;|&check;|285|Da||
|4.1.4|[DELETED, DUPLICATE OF 4.1.3]|||||||
|4.1.5|Verify that access controls fail securely including when an exception occurs. (C10)|&check;|&check;|&check;|285|Da|Pad servisa za access control neće dozvoliti nikom da pristup dokumentima|


### V4.2 Operation Level Access Control

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|4.2.1|Verify that sensitive data and APIs are protected against Insecure Direct Object Reference (IDOR) attacks targeting creation, reading, updating and deletion of records, such as creating or updating someone else's record, viewing everyone's records, or deleting all records.|&check;|&check;|&check;|639|Da||
|4.2.2|Verify that the application or framework enforces a strong anti-CSRF mechanism to protect authenticated functionality, and effective antiautomation or anti-CSRF protects unauthenticated functionality.|&check;|&check;|&check;|352|Da||


### V4.3 Other Access Control Considerations

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|4.3.1|Verify administrative interfaces use appropriate multi-factor authentication to prevent unauthorized use.|&check;|&check;|&check;|419|Ne|Nemamo multi-factor autentifikaciju|
|4.3.2|Verify that directory browsing is disabled unless deliberately desired. Additionally, applications should not allow discovery or disclosure of file or directory metadata, such as Thumbs.db, .DS_Store, .git or .svn folders.|&check;|&check;|&check;|548|Da|Korisnici ne mogu da pretražuju direktorijume|
|4.3.3|Verify the application has additional authorization (such as step up or adaptive authentication) for lower value systems, and / or segregation of duties for high value applications to enforce anti-fraud controls as per the risk of application and past fraud.||&check;|&check;|732|Ne||


## V5 Validation, Sanitization and Encoding

### V5.1 Input Validation

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|5.1.1|Verify that the application has defenses against HTTP parameter pollution attacks, particularly if the application framework makes no distinction about the source of request parameters (GET, POST, cookies, headers, or environment variables).|&check;|&check;|&check;|235|Da||
|5.1.2|Verify that frameworks protect against mass parameter assignment attacks, or that the application has countermeasures to protect against unsafe parameter assignment, such as marking fields private or similar. (C5)|&check;|&check;|&check;|915|Da||
|5.1.3|Verify that all input (HTML form fields, REST requests, URL parameters, HTTP headers, cookies, batch files, RSS feeds, etc) is validated using positive validation (allow lists). (C5)|&check;|&check;|&check;|20|Da||
|5.1.4|Verify that structured data is strongly typed and validated against a defined schema including allowed characters, length and pattern (e.g. credit card numbers, e-mail addresses, telephone numbers, or validating that two related fields are reasonable, such as checking that suburb and zip/postcode match). (C5)|&check;|&check;|&check;|20|Da||
|5.1.5|Verify that URL redirects and forwards only allow destinations which appear on an allow list, or show a warning when redirecting to potentially untrusted content.|&check;|&check;|&check;|601|Da||


### V5.2 Sanitization and Sandboxing

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|5.2.1|Verify that all untrusted HTML input from WYSIWYG editors or similar is properly sanitized with an HTML sanitizer library or framework feature. (C5)|&check;|&check;|&check;|116|Da|Nemamo HTML input|
|5.2.2|Verify that unstructured data is sanitized to enforce safety measures such as allowed characters and length.|&check;|&check;|&check;|138|Da||
|5.2.3|Verify that the application sanitizes user input before passing to mail systems to protect against SMTP or IMAP injection.|&check;|&check;|&check;|147|Da|Ne šaljemo mejlove|
|5.2.4|Verify that the application avoids the use of eval() or other dynamic code execution features. Where there is no alternative, any user input being included must be sanitized or sandboxed before being executed.|&check;|&check;|&check;|95|Da|Ne koristimo eval i slicno|
|5.2.5|Verify that the application protects against template injection attacks by ensuring that any user input being included is sanitized or sandboxed.|&check;|&check;|&check;|94|Da||
|5.2.6|Verify that the application protects against SSRF attacks, by validating or sanitizing untrusted data or HTTP file metadata, such as filenames and URL input fields, and uses allow lists of protocols, domains, paths and ports.|&check;|&check;|&check;|918|Ne||
|5.2.7|Verify that the application sanitizes, disables, or sandboxes user-supplied Scalable Vector Graphics (SVG) scriptable content, especially as they relate to XSS resulting from inline scripts, and foreignObject.|&check;|&check;|&check;|159|Da|Ne koristimo SVG|
|5.2.8|Verify that the application sanitizes, disables, or sandboxes user-supplied scriptable or expression template language content, such as Markdown, CSS or XSL stylesheets, BBCode, or similar.|&check;|&check;|&check;|94|Da||


### V5.3 Output Encoding and Injection Prevention

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|5.3.1|Verify that output encoding is relevant for the interpreter and context required. For example, use encoders specifically for HTML values, HTML attributes, JavaScript, URL parameters, HTTP headers, SMTP, and others as the context requires, especially from untrusted inputs (e.g. names with Unicode or apostrophes, such as ねこ or O'Hara). (C4)|&check;|&check;|&check;|116|Da||
|5.3.2|Verify that output encoding preserves the user's chosen character set and locale, such that any Unicode character point is valid and safely handled. (C4)|&check;|&check;|&check;|176|Da||
|5.3.3|Verify that context-aware, preferably automated - or at worst, manual - output escaping protects against reflected, stored, and DOM based XSS. (C4)|&check;|&check;|&check;|79|Da||
|5.3.4|Verify that data selection or database queries (e.g. SQL, HQL, ORM, NoSQL) use parameterized queries, ORMs, entity frameworks, or are otherwise protected from database injection attacks. (C3)|&check;|&check;|&check;|89|Da|Koristimo prepared statement|
|5.3.5|Verify that where parameterized or safer mechanisms are not present, context-specific output encoding is used to protect against injection attacks, such as the use of SQL escaping to protect against SQL injection. (C3, C4)|&check;|&check;|&check;|89|Da||
|5.3.6|Verify that the application protects against JSON injection attacks, JSON eval attacks, and JavaScript expression evaluation. (C4)|&check;|&check;|&check;|830|Da||
|5.3.7|Verify that the application protects against LDAP injection vulnerabilities, or that specific security controls to prevent LDAP injection have been implemented. (C4)|&check;|&check;|&check;|90|Ne|Nismo razmatrali ovu ranjivost|
|5.3.8|Verify that the application protects against OS command injection and that operating system calls use parameterized OS queries or use contextual command line output encoding. (C4)|&check;|&check;|&check;|78|Da|Ne koristimo OS komande|
|5.3.9|Verify that the application protects against Local File Inclusion (LFI) or Remote File Inclusion (RFI) attacks.|&check;|&check;|&check;|829|Da||
|5.3.10|Verify that the application protects against XPath injection or XML injection attacks. (C4)|&check;|&check;|&check;|643|Da|Ne koristimo XML i XPath|


### V5.4 Memory, String, and Unmanaged Code

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|5.4.1|Verify that the application uses memory-safe string, safer memory copy and pointer arithmetic to detect or prevent stack, buffer, or heap overflows.||&check;|&check;|120|Da||
|5.4.2|Verify that format strings do not take potentially hostile input, and are constant.||&check;|&check;|134|Da||
|5.4.3|Verify that sign, range, and input validation techniques are used to prevent integer overflows.||&check;|&check;|190|Da||


### V5.5 Deserialization Prevention

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|5.5.1|Verify that serialized objects use integrity checks or are encrypted to prevent hostile object creation or data tampering. (C5)|&check;|&check;|&check;|502|Da||
|5.5.2|Verify that the application correctly restricts XML parsers to only use the most restrictive configuration possible and to ensure that unsafe features such as resolving external entities are disabled to prevent XML eXternal Entity (XXE) attacks.|&check;|&check;|&check;|611|Da|Ne koristimo XML parser|
|5.5.3|Verify that deserialization of untrusted data is avoided or is protected in both custom code and third-party libraries (such as JSON, XML and YAML parsers).|&check;|&check;|&check;|502|Da||
|5.5.4|Verify that when parsing JSON in browsers or JavaScript-based backends, JSON.parse is used to parse the JSON document. Do not use eval() to parse JSON.|&check;|&check;|&check;|95|Da|Ne koristimo eval|


## V6 Stored Cryptography

### V6.1 Data Classification

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|6.1.1|Verify that regulated private data is stored encrypted while at rest, such as Personally Identifiable Information (PII), sensitive personal information, or data assessed likely to be subject to EU's GDPR.||&check;|&check;|311|Ne|Ne čuvamo takve informacije|
|6.1.2|Verify that regulated health data is stored encrypted while at rest, such as medical records, medical device details, or de-anonymized research records.||&check;|&check;|311|Ne||
|6.1.3|Verify that regulated financial data is stored encrypted while at rest, such as financial accounts, defaults or credit history, tax records, pay history, beneficiaries, or de-anonymized market or research records.||&check;|&check;|311|Ne||


### V6.2 Algorithms

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|6.2.1|Verify that all cryptographic modules fail securely, and errors are handled in a way that does not enable Padding Oracle attacks.|&check;|&check;|&check;|310|Da||
|6.2.2|Verify that industry proven or government approved cryptographic algorithms, modes, and libraries are used, instead of custom coded cryptography. (C8)||&check;|&check;|327|Da||
|6.2.3|Verify that encryption initialization vector, cipher configuration, and block modes are configured securely using the latest advice.||&check;|&check;|326|Ne||
|6.2.4|Verify that random number, encryption or hashing algorithms, key lengths, rounds, ciphers or modes, can be reconfigured, upgraded, or swapped at any time, to protect against cryptographic breaks. (C8)||&check;|&check;|326|Ne|Nije lako zameniti algoritam|
|6.2.5|Verify that known insecure block modes (i.e. ECB, etc.), padding modes (i.e. PKCS#1 v1.5, etc.), ciphers with small block sizes (i.e. Triple-DES, Blowfish, etc.), and weak hashing algorithms (i.e. MD5, SHA1, etc.) are not used unless required for backwards compatibility.||&check;|&check;|326|Da||
|6.2.6|Verify that nonces, initialization vectors, and other single use numbers must not be used more than once with a given encryption key. The method of generation must be appropriate for the algorithm being used.||&check;|&check;|326|Da||
|6.2.7|Verify that encrypted data is authenticated via signatures, authenticated cipher modes, or HMAC to ensure that ciphertext is not altered by an unauthorized party.|||&check;|326|Da|JWT tokeni su potpisani|
|6.2.8|Verify that all cryptographic operations are constant-time, with no 'short-circuit' operations in comparisons, calculations, or returns, to avoid leaking information.|||&check;|385|Da||

Trebalo bi omogućiti laku zamenu kriptografskih algoritama.


### V6.3 Random Values

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|6.3.1|Verify that all random numbers, random file names, random GUIDs, and random strings are generated using the cryptographic module's approved cryptographically secure random number generator when these random values are intended to be not guessable by an attacker.||&check;|&check;|338|Da|Kriptografske biblioteke koriste secure random number generator|
|6.3.2|Verify that random GUIDs are created using the GUID v4 algorithm, and a Cryptographically-secure Pseudo-random Number Generator (CSPRNG). GUIDs created using other pseudo-random number generators may be predictable.||&check;|&check;|338|Ne|Ne generišemo GUID|
|6.3.3|Verify that random numbers are created with proper entropy even when the application is under heavy load, or that the application degrades gracefully in such circumstances.|||&check;|338|Ne||


### V6.4 Secret Management

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|6.4.1|Verify that a secrets management solution such as a key vault is used to securely create, store, control access to and destroy secrets. (C8)||&check;|&check;|798|Ne|Nismo koristili key vault|
|6.4.2|Verify that key material is not exposed to the application but instead uses an isolated security module like a vault for cryptographic operations. (C8)||&check;|&check;|320|Ne||


## V7 Error Handling and Logging

### V7.1 Log Content

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|7.1.1|Verify that the application does not log credentials or payment details. Session tokens should only be stored in logs in an irreversible, hashed form. (C9, C10)|&check;|&check;|&check;|532|Da||
|7.1.2|Verify that the application does not log other sensitive data as defined under local privacy laws or relevant security policy. (C9)|&check;|&check;|&check;|532|Da||
|7.1.3|Verify that the application logs security relevant events including successful and failed authentication events, access control failures, deserialization failures and input validation failures. (C5, C7)||&check;|&check;|778|Ne||
|7.1.4|Verify that each log event includes necessary information that would allow for a detailed investigation of the timeline when an event happens. (C9)||&check;|&check;|778|Ne||

Trebalo bi da logujemo više bitnih događaja i sa više detalja kao što je vreme kada se događaj desio.


### V7.2 Log Processing

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|7.2.1|Verify that all authentication decisions are logged, without storing sensitive session tokens or passwords. This should include requests with relevant metadata needed for security investigations.||&check;|&check;|778|Ne||
|7.2.2|Verify that all access control decisions can be logged and all failed decisions are logged. This should include requests with relevant metadata needed for security investigations.||&check;|&check;|285|Ne||

Trebalo bi da logujemo sa dovoljno detalja da je moguće raditi monitoring logova.


### V7.3 Log Protection

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|7.3.1|Verify that all logging components appropriately encode data to prevent log injection. (C9)||&check;|&check;|117|Da||
|7.3.2|[DELETED, DUPLICATE OF 7.3.1]||||||
|7.3.3|Verify that security logs are protected from unauthorized access and modification. (C9)||&check;|&check;|200|Da||
|7.3.4|Verify that time sources are synchronized to the correct time and time zone. Strongly consider logging only in UTC if systems are global to assist with postincident forensic analysis. (C9)||&check;|&check;||Ne||

Trebalo bi da logujemo i vremena u UTC.


### V7.4 Error Handling

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|7.4.1|Verify that a generic message is shown when an unexpected or security sensitive error occurs, potentially with a unique ID which support personnel can use to investigate. (C10)|&check;|&check;|&check;|210|Ne|Nisu sve greske logovane|
|7.4.2|Verify that exception handling (or a functional equivalent) is used across the codebase to account for expected and unexpected error conditions. (C10)||&check;|&check;|544|Da|Funkcije vracaju error kao drugu povratnu vrednost|
|7.4.3|Verify that a "last resort" error handler is defined which will catch all unhandled exceptions. (C10)||&check;|&check;|431|Ne||


## V8 Data Protection

### V8.1 General Data Protection

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|8.1.1|Verify the application protects sensitive data from being cached in server components such as load balancers and application caches.||&check;|&check;|524|Da|Ne koristimo load balancer|
|8.1.2|Verify that all cached or temporary copies of sensitive data stored on the server are protected from unauthorized access or purged/invalidated after the authorized user accesses the sensitive data.||&check;|&check;|524|Da||
|8.1.3|Verify the application minimizes the number of parameters in a request, such as hidden fields, Ajax variables, cookies and header values.||&check;|&check;|233|Da|Broj parametara je minimalan potreban|
|8.1.4|Verify the application can detect and alert on abnormal numbers of requests, such as by IP, user, total per hour or day, or whatever makes sense for the application.||&check;|&check;|770|Ne||
|8.1.5|Verify that regular backups of important data are performed and that test restoration of data is performed.|||&check;|19|Ne|Ne pravimo backup|
|8.1.6|Verify that backups are stored securely to prevent data from being stolen or corrupted.|||&check;|19|Ne||

Trebalo bi da pravimo backup regularno i da iskoristimo npr cloud flare za sprečavanje abnormalnog broja zahteva od istog korisnika.


### V8.2 Client-side Data Protection

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|8.2.1|Verify the application sets sufficient anti-caching headers so that sensitive data is not cached in modern browsers.|&check;|&check;|&check;|525|Ne|Nismo koristili anti-cashing hedere|
|8.2.2|Verify that data stored in browser storage (such as localStorage, sessionStorage, IndexedDB, or cookies) does not contain sensitive data.|&check;|&check;|&check;|922|Ne|Čuvamo JWT token u localStorage|
|8.2.3|Verify that authenticated data is cleared from client storage, such as the browser DOM, after the client or session is terminated.|&check;|&check;|&check;|922|Ne||

Trebalo bi da sprečimo keširanje osetljivih podataka.


### V8.3 Sensitive Private Data

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|8.3.1|Verify that sensitive data is sent to the server in the HTTP message body or headers, and that query string parameters from any HTTP verb do not contain sensitive data.|&check;|&check;|&check;|319|Da|Query parametri se ne koriste za osetljive podatke|
|8.3.2|Verify that users have a method to remove or export their data on demand.|&check;|&check;|&check;|212|Ne|Nemamo export podataka, a brisanje je moguće za neke podatke|
|8.3.3|Verify that users are provided clear language regarding collection and use of supplied personal information and that users have provided opt-in consent for the use of that data before it is used in any way.|&check;|&check;|&check;|285|Ne|Korisnici nisu obavesteni o čuvanju njihovih podataka|
|8.3.4|Verify that all sensitive data created and processed by the application has been identified, and ensure that a policy is in place on how to deal with sensitive data. (C8)|&check;|&check;|&check;|200|Ne||
|8.3.5|Verify accessing sensitive data is audited (without logging the sensitive data itself), if the data is collected under relevant data protection directives or where logging of access is required.||&check;|&check;|532|Ne|Nemamo dovoljno logova|
|8.3.6|Verify that sensitive information contained in memory is overwritten as soon as it is no longer required to mitigate memory dumping attacks, using zeroes or random data.||&check;|&check;|226|Ne||
|8.3.7|Verify that sensitive or private information that is required to be encrypted, is encrypted using approved algorithms that provide both confidentiality and integrity. (C8)||&check;|&check;|327|Da||
|8.3.8|Verify that sensitive personal information is subject to data retention classification, such that old or out of date data is deleted automatically, on a schedule, or as the situation requires.||&check;|&check;|285|Ne|Ne brišemo stare podatke automatski|

Trebalo bi da uskladimo naše čuvanje podataka sa GDPR.


## V9 Communication

### V9.1 Client Communication Security

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|9.1.1|Verify that TLS is used for all client connectivity, and does not fall back to insecure or unencrypted communications. (C8)|&check;|&check;|&check;|319|Da|Koristimo https za svu komunikaciju|
|9.1.2|Verify using up to date TLS testing tools that only strong cipher suites are enabled, with the strongest cipher suites set as preferred.|&check;|&check;|&check;|326|Ne|Nemamo TLS testing tool|
|9.1.3|Verify that only the latest recommended versions of the TLS protocol are enabled, such as TLS 1.2 and TLS 1.3. The latest version of the TLS protocol should be the preferred option.|&check;|&check;|&check;|326|Da||


### V9.2 Server Communication Security

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|9.2.1|Verify that connections to and from the server use trusted TLS certificates. Where internally generated or self-signed certificates are used, the server must be configured to only trust specific internal CAs and specific self-signed certificates. All others should be rejected.||&check;|&check;|295|Da||
|9.2.2|Verify that encrypted communications such as TLS is used for all inbound and outbound connections, including for management ports, monitoring, authentication, API, or web service calls, database, cloud, serverless, mainframe, external, and partner connections. The server must not fall back to insecure or unencrypted protocols.||&check;|&check;|319|Ne|Nemamo enkriptovanu komunikaciju sa bazom|
|9.2.3|Verify that all encrypted connections to external systems that involve sensitive information or functions are authenticated.||&check;|&check;|287|Da||
|9.2.4|Verify that proper certification revocation, such as Online Certificate Status Protocol (OCSP) Stapling, is enabled and configured.||&check;|&check;|299|Ne|Nismo ovo konfigurisali|
|9.2.5|Verify that backend TLS connection failures are logged.|||&check;|544|Ne|Ne logujemo TLS connection failure|

Trebalo bi da enkriptujemo komunikaciju sa bazom i da omogućimo povlačenje sertifikata.


## V10 Malicious Code

### V10.1 Code Integrity

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|10.1.1|Verify that a code analysis tool is in use that can detect potentially malicious code, such as time functions, unsafe file operations and network connections.|||&check;|749|Da|Koristimo CodeQL|


### V10.2 Malicious Code Search

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|10.2.1|Verify that the application source code and third party libraries do not contain unauthorized phone home or data collection capabilities. Where such functionality exists, obtain the user's permission for it to operate before collecting any data.||&check;|&check;|359|Da||
|10.2.2|Verify that the application does not ask for unnecessary or excessive permissions to privacy related features or sensors, such as contacts, cameras, microphones, or location.||&check;|&check;|272|Da||
|10.2.3|Verify that the application source code and third party libraries do not contain back doors, such as hard-coded or additional undocumented accounts or keys, code obfuscation, undocumented binary blobs, rootkits, or anti-debugging, insecure debugging features, or otherwise out of date, insecure, or hidden functionality that could be used maliciously if discovered.|||&check;|507|Da||
|10.2.4|Verify that the application source code and third party libraries do not contain time bombs by searching for date and time related functions.|||&check;|511|Da||
|10.2.5|Verify that the application source code and third party libraries do not contain malicious code, such as salami attacks, logic bypasses, or logic bombs.|||&check;|511|Da||
|10.2.6|Verify that the application source code and third party libraries do not contain Easter eggs or any other potentially unwanted functionality.|||&check;|507|Da||


### V10.3 Application Integrity

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|10.3.1|Verify that if the application has a client or server auto-update feature, updates should be obtained over secure channels and digitally signed. The update code must validate the digital signature of the update before installing or executing the update.|&check;|&check;|&check;|16|Da|Nemamo automatski update|
|10.3.2|Verify that the application employs integrity protections, such as code signing or subresource integrity. The application must not load or execute code from untrusted sources, such as loading includes, modules, plugins, code, or libraries from untrusted sources or the Internet.|&check;|&check;|&check;|353|Ne|Ne potpisujemo kod|
|10.3.3|Verify that the application has protection from subdomain takeovers if the application relies upon DNS entries or DNS subdomains, such as expired domain names, out of date DNS pointers or CNAMEs, expired projects at public source code repos, or transient cloud APIs, serverless functions, or storage buckets (autogen-bucket-id.cloud.example.com) or similar. Protections can include ensuring that DNS names used by applications are regularly checked for expiry or change.|&check;|&check;|&check;|350|Ne||


## V11 Business Logic

### V11.1 Business Logic Security

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|11.1.1|Verify that the application will only process business logic flows for the same user in sequential step order and without skipping steps|&check;|&check;|&check;|841|Da||
|11.1.2|Verify that the application will only process business logic flows with all steps being processed in realistic human time, i.e. transactions are not submitted too quickly.|&check;|&check;|&check;|799|Ne|Nemamo proveru da li se akcije odvijaju previše brzo|
|11.1.3|Verify the application has appropriate limits for specific business actions or transactions which are correctly enforced on a per user basis.|&check;|&check;|&check;|770|Ne||
|11.1.4|Verify that the application has anti-automation controls to protect against excessive calls such as mass data exfiltration, business logic requests, file uploads or denial of service attacks.|&check;|&check;|&check;|770|Ne||
|11.1.5|Verify the application has business logic limits or validation to protect against likely business risks or threats, identified using threat modeling or similar methodologies.|&check;|&check;|&check;|841|Da|Napravili smo threat model|
|11.1.6|Verify that the application does not suffer from "Time Of Check to Time Of Use" (TOCTOU) issues or other race conditions for sensitive operations.||&check;|&check;|367|Da||
|11.1.7|Verify that the application monitors for unusual events or activity from a business logic perspective. For example, attempts to perform actions out of order or actions which a normal user would never attempt. (C9)||&check;|&check;|754|Ne||
|11.1.8|Verify that the application has configurable alerting when automated attacks or unusual activity is detected.||&check;|&check;|390|Ne||

Trebalo bi raditi monitoring tako da možemo da detektujemo kada se akcije ivršavaju previše brzo ili u pogrešnom redosledu.


## V12 Files and Resources

### V12.1 File Upload

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|12.1.1|Verify that the application will not accept large files that could fill up storage or cause a denial of service.|&check;|&check;|&check;|400|Ne|Nismo ograničili veličine fajlova|
|12.1.2|Verify that the application checks compressed files (e.g. zip, gz, docx, odt) against maximum allowed uncompressed size and against maximum number of files before uncompressing the file.||&check;|&check;|409|Ne|Nemamo upload kompresovanih fajlova|
|12.1.3|Verify that a file size quota and maximum number of files per user is enforced to ensure that a single user cannot fill up the storage with too many files, or excessively large files.||&check;|&check;|770|Ne||

Trebalo bi ograničiti broj i veličinu fajlova koji mogu da se uploaduju.


### V12.2 File Integrity

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|12.2.1|Verify that files obtained from untrusted sources are validated to be of expected type based on the file's content.||&check;|&check;|434|Da||


### V12.3 File Execution

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|12.3.1|Verify that user-submitted filename metadata is not used directly by system or framework filesystems and that a URL API is used to protect against path traversal.|&check;|&check;|&check;|22|Da||
|12.3.2|Verify that user-submitted filename metadata is validated or ignored to prevent the disclosure, creation, updating or removal of local files (LFI).|&check;|&check;|&check;|73|Ne||
|12.3.3|Verify that user-submitted filename metadata is validated or ignored to prevent the disclosure or execution of remote files via Remote File Inclusion (RFI) or Server-side Request Forgery (SSRF) attacks.|&check;|&check;|&check;|98|Ne||
|12.3.4|Verify that the application protects against Reflective File Download (RFD) by validating or ignoring user-submitted filenames in a JSON, JSONP, or URL parameter, the response Content-Type header should be set to text/plain, and the Content-Disposition header should have a fixed filename.|&check;|&check;|&check;|641|Da||
|12.3.5|Verify that untrusted file metadata is not used directly with system API or libraries, to protect against OS command injection.|&check;|&check;|&check;|78|Ne||
|12.3.6|Verify that the application does not include and execute functionality from untrusted sources, such as unverified content distribution networks, JavaScript libraries, node npm libraries, or server-side DLLs.||&check;|&check;|829|Da||

Trebalo bi obrisati metapodatne iz fajlova koji se uploaduju.


### V12.4 File Storage

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|12.4.1|Verify that files obtained from untrusted sources are stored outside the web root, with limited permissions.|&check;|&check;|&check;|552|Da||
|12.4.2|Verify that files obtained from untrusted sources are scanned by antivirus scanners to prevent upload and serving of known malicious content.|&check;|&check;|&check;|509|Ne|Nemamo antivirus|


### V12.5 File Download

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|12.5.1|Verify that the web tier is configured to serve only files with specific file extensions to prevent unintentional information and source code leakage. For example, backup files (e.g. .bak), temporary working files (e.g. .swp), compressed files (.zip, .tar.gz, etc) and other extensions commonly used by editors should be blocked unless required.|&check;|&check;|&check;|552|Da||
|12.5.2|Verify that direct requests to uploaded files will never be executed as HTML/JavaScript content.|&check;|&check;|&check;|434|Da||

### V12.6 SSRF Protection

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|12.6.1|Verify that the web or application server is configured with an allow list of resources or systems to which the server can send requests or load data/files from.|&check;|&check;|&check;|918|Ne||


## V13 API and Web Service

### V13.1 Generic Web Service Security

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|13.1.1|Verify that all application components use the same encodings and parsers to avoid parsing attacks that exploit different URI or file parsing behavior that could be used in SSRF and RFI attacks.|&check;|&check;|&check;|116|Da||
|13.1.2|[DELETED, DUPLICATE OF 4.3.1]||||||
|13.1.3|Verify API URLs do not expose sensitive information, such as the API key, session tokens etc.|&check;|&check;|&check;|598|Da||
|13.1.4|Verify that authorization decisions are made at both the URI, enforced by programmatic or declarative security at the controller or router, and at the resource level, enforced by model-based permissions.||&check;|&check;|285|Da||
|13.1.5|Verify that requests containing unexpected or missing content types are rejected with appropriate headers (HTTP response status 406 Unacceptable or 415 Unsupported Media Type).||&check;|&check;|434|Da||


### V13.2 RESTful Web Service

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|13.2.1|Verify that enabled RESTful HTTP methods are a valid choice for the user or action, such as preventing normal users using DELETE or PUT on protected API or resources.|&check;|&check;|&check;|650|Da||
|13.2.2|Verify that JSON schema validation is in place and verified before accepting input.|&check;|&check;|&check;|20|Da||
|13.2.3|Verify that RESTful web services that utilize cookies are protected from Cross-Site Request Forgery via the use of at least one or more of the following: double submit cookie pattern, CSRF nonces, or Origin request header checks.|&check;|&check;|&check;|352|Da|Ne koristimo cookies|
|13.2.4|[DELETED, DUPLICATE OF 11.1.4]||||||
|13.2.5|Verify that REST services explicitly check the incoming Content-Type to be the expected one, such as application/xml or application/json.||&check;|&check;|436|Da||
|13.2.6|Verify that the message headers and payload are trustworthy and not modified in transit. Requiring strong encryption for transport (TLS only) may be sufficient in many cases as it provides both confidentiality and integrity protection. Per-message digital signatures can provide additional assurance on top of the transport protections for high-security applications but bring with them additional complexity and risks to weigh against the benefits.||&check;|&check;|345|Da|Koristimo https|


### V13.3 SOAP Web Service

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|13.3.1|Verify that XSD schema validation takes place to ensure a properly formed XML document, followed by validation of each input field before any processing of that data takes place.|&check;|&check;|&check;|20|Da|Ne koristimo SOAP|
|13.3.2|Verify that XSD schema validation takes place to ensure a properly formed XML document, followed by validation of each input field before any processing of that data takes place.||&check;|&check;|345|Da|Ne koristimo SOAP|


### V13.4 GraphQL

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|13.4.1|Verify that a query allow list or a combination of depth limiting and amount limiting is used to prevent GraphQL or data layer expression Denial of Service (DoS) as a result of expensive, nested queries. For more advanced scenarios, query cost analysis should be used.||&check;|&check;|770|Da|Ne koristimo GraphQL|
|13.4.2|Verify that GraphQL or other data layer authorization logic should be implemented at the business logic layer instead of the GraphQL layer||&check;|&check;|285|Da|Ne koristimo GraphQL|


## V14 Configuration

### V14.1 Build and Deploy

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|14.1.1|Verify that the application build and deployment processes are performed in a secure and repeatable way, such as CI / CD automation, automated configuration management, and automated deployment scripts.||&check;|&check;||Da||
|14.1.2|Verify that compiler flags are configured to enable all available buffer overflow protections and warnings, including stack randomization, data execution prevention, and to break the build if an unsafe pointer, memory, format string, integer, or string operations are found.||&check;|&check;|120|Da||
|14.1.3|Verify that server configuration is hardened as per the recommendations of the application server and frameworks in use.||&check;|&check;|16|Da||
|14.1.4|Verify that the application, configuration, and all dependencies can be redeployed using automated deployment scripts, built from a documented and tested runbook in a reasonable time, or restored from backups in a timely fashion.||&check;|&check;||Ne|Nismo deployovali aplikaciju|
|14.1.5|Verify that authorized administrators can verify the integrity of all securityrelevant configurations to detect tampering.|||&check;||Da||


### V14.2 Dependency

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|14.2.1|Verify that all components are up to date, preferably using a dependency checker during build or compile time. (C2)|&check;|&check;|&check;|1026|Da||
|14.2.2|Verify that all unneeded features, documentation, sample applications and configurations are removed.|&check;|&check;|&check;|1002|Da||
|14.2.3|Verify that if application assets, such as JavaScript libraries, CSS or web fonts, are hosted externally on a Content Delivery Network (CDN) or external provider, Subresource Integrity (SRI) is used to validate the integrity of the asset.|&check;|&check;|&check;|829|Ne||
|14.2.4|Verify that third party components come from pre-defined, trusted and continually maintained repositories. (C2)||&check;|&check;|829|Da||
|14.2.5|Verify that a Software Bill of Materials (SBOM) is maintained of all third party libraries in use. (C2)||&check;|&check;||Ne||
|14.2.6|Verify that the attack surface is reduced by sandboxing or encapsulating third party libraries to expose only the required behaviour into the application. (C2)||&check;|&check;|265|Ne||


### V14.3 Unintended Security Disclosure

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|14.3.1|[DELETED, DUPLICATE OF 7.4.1]||||||
|14.3.2|Verify that web or application server and application framework debug modes are disabled in production to eliminate debug features, developer consoles, and unintended security disclosures.|&check;|&check;|&check;|497|Da||
|14.3.3|Verify that the HTTP headers or any part of the HTTP response do not expose detailed version information of system components.|&check;|&check;|&check;|200|Da||


### V14.4 HTTP Security Headers

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|14.4.1|Verify that every HTTP response contains a Content-Type header. Also specify a safe character set (e.g., UTF-8, ISO-8859-1) if the content types are text/*, /+xml and application/xml. Content must match with the provided Content-Type header.|&check;|&check;|&check;|173|Da||
|14.4.2|Verify that all API responses contain a Content-Disposition: attachment; filename="api.json" header (or other appropriate filename for the content type).|&check;|&check;|&check;|116|Ne||
|14.4.3|Verify that a Content Security Policy (CSP) response header is in place that helps mitigate impact for XSS attacks like HTML, DOM, JSON, and JavaScript injection vulnerabilities.|&check;|&check;|&check;|1021|Da||
|14.4.4|Verify that all responses contain a X-Content-Type-Options: nosniff header.|&check;|&check;|&check;|116|Da||
|14.4.5|Verify that a Strict-Transport-Security header is included on all responses and for all subdomains, such as Strict-Transport-Security: maxage=15724800; includeSubdomains.|&check;|&check;|&check;|523|Ne||
|14.4.6|Verify that a suitable Referrer-Policy header is included to avoid exposing sensitive information in the URL through the Referer header to untrusted parties.|&check;|&check;|&check;|116|Ne||
|14.4.7|Verify that the content of a web application cannot be embedded in a thirdparty site by default and that embedding of the exact resources is only allowed where necessary by using suitable Content-Security-Policy: frameancestors and X-Frame-Options response headers.|&check;|&check;|&check;|1021|Da||


### V14.5 HTTP Request Header Validation

|Identifikator|Opis|L1|L2|L3|CWE|Ispunjeno?||
|---|---|---|---|---|---|---|---|
|14.5.1|Verify that the application server only accepts the HTTP methods in use by the application/API, including pre-flight OPTIONS, and logs/alerts on any requests that are not valid for the application context.|&check;|&check;|&check;|749|Da||
|14.5.2|Verify that the supplied Origin header is not used for authentication or access control decisions, as the Origin header can easily be changed by an attacker.|&check;|&check;|&check;|346|Da||
|14.5.3|Verify that the Cross-Origin Resource Sharing (CORS) Access-Control-AllowOrigin header uses a strict allow list of trusted domains and subdomains to match against and does not support the "null" origin.|&check;|&check;|&check;|346|Ne||
|14.5.4|Verify that HTTP headers added by a trusted proxy or SSO devices, such as a bearer token, are authenticated by the application.||&check;|&check;|306|Da||

