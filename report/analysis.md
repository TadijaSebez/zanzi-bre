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

