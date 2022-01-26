// TCP server som hører på innkommende TCP forbindelser,
// initialiserer forbindelser med klienten, aksepterer/godtar
// og asynkront behandler hver forbindelse, utveksler data
// og avslutter forbindelse.
package details

import (
    "net"
    "testing"
)

/* 
   net.Listen funksjon aksepterer nettverkstype (tcp) og
   en IP adresse, samt en portnummer separert med kolon.
   Funksjonen returnerer en net.Listener grensesnitt og 
   en "error" grensesnitt. (kan også gi tcp4 eller tcp6)
   som det første argumentet.

   Hvis funksjonen utfører korrekt, blir "listener" bundet
   til en IP adresse og portnummer. 

   Bindingen betyr at _OS_ har eksklusivt gitt porten og 
   den gitte IP adressen til tilhører ("listener"). 
   
   Andre prosesser kan ikke bruke samme port. 
   
   Hvis portnummer er 0 eller tom, Go will tildele et 
   tilfeldig portnummer til tilhøreren ("listener").
   
   "listener" sin adresse kan man oppnå med et kall til Addr 

   Hvis IP adresse er ikke spesifisert, "listener" vil bli 
   bundet til alle "unicast" og "anycast" IP adresser på 
   systemet.
   
   Hvis både IP adresse og port ikke er spesifisert (eller 
   det er lagt inn ":" som det andre argumentet til net.Listen 
   vil "listener" bindes til alle "unicast" og "anycast" IP
   adresser og bruke et tilfeldig portnummer. 
    
   Viktig å avslutte (frakoble) "listener". Det er når 
   "listener" sin Accept funksjon blir påkalt, at den 
   kan blokkere/vente i ubestemt tid. Når man bruker 
   Close, avblokkeres alle kall til Accept funksjonen. 
*/
func TestListener(t *testing.T) {
    listener, err := net.Listen("tcp", "127.0.0.1:0")
    if err != nil {
        t.Fatal(err)
    }

    defer func() { _ = listener.Close() }()
    t.Logf("bound to %q", listener.Addr())
}
