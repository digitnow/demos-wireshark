# TCP
Formålet er å feilsøke nettverksproblemer i Go programmer.

## TCP aspekter:
- håndtrykk
- sekvensnummere
- bekreftelse
- retransmisjon (ikke videresending, men sending på nytt?)
- 

## TCP prosess
- dialing (oppringning)
- listening (høring)
- accepting (godkjennelse)
- termination (avslutning)

## TCP other issues
- time-outs
- temporary errors (hvordan oppdage disse? hvordan gjøre det brukbart?)
- detection of unreliable network connections

## Hvorfor er TCP en pålitelig protokoll?
- løser problemer med pakketap (pga. interferens eller overbelastning av nettverk / "trafikkork" når noder sender mer data inn i nettverket enn forbindelsen kan behandle og overflødige pakker blir kastet)
- TCP implementasjoner holder datatap til et minimum, selv når nettverkstilstand endres, - det vil si de inneholder flytkontroll
- løser problemer med mottak av pakker med ødelagt sekvens (istedenfor "universitetet" mottaker mottar "unievrsiettet")

## TCP sesjon
- abstraksjon lar programmerer å levere en strøm av data til mottaker og motta en bekreftelse om at mottaker har mottatt data (unngå å sende mye data inn i nettverket selv om mottaker "sover" eller kan ikke motta data)
- eksempel kan være en dialog mellom to mennesker hvor den som hører kontinuerlig bekrefter til den som snakker om at informasjon er mottatt; det gjør det mulig å åndtere feil i sanntid 
- generell samtale mellom to noder, - hilser først, utveklser noen meldinger og sier ha det
- sesjon forblir tomhendt inntil en av partene har noe å si (har data for å utveksle); TCP sesjonene som ikke behandles korrekt kan resultere i unødvendig minneforbruk (man kan bruke "time-outs" for å unngå det)


## Håndtrykk
- Tilstander klient: Dial, Established
- Tilstander server: Listen, Accept, Established
- Bruke "klient" og "server" for å beskrive noden som ringer og noden som hører på innringinger

SYN -> (informerer server om sine evner og hvilke vindusstørrelse foretrekkes)
  
   (responderer til oppringeren med både ACK og SYN biter satt)  <- SYN/ACK 

ACK -> (svarer til serveren med ACK pakke for bekrefte serverens SYN pakke)



## grensesnitt i "net" pakke



## Eksempler
- Sender data med en rate på 1 Gigabit per sekundet over et nettverk som kan behandle 10 Megabit per sekundet (nettverket blir mettet og kaster overflødige data)
- "Intermediate nodes, such as firewalls, can send RST packets to each node in a connection, effectively terminating the socket from the middle." er det det som skjer i Win10 med Golang TCP klient?
