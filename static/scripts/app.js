//
var A=0; window.db=null;

//
function ajax(u, cb){
  var x = new XMLHttpRequest();
  x.open('GET', u, false); 
  try{ x.send(); }catch(e){ console.log('oups',e); }
  if(x.status==200){ cb(JSON.parse(x.responseText)); } else { cb([]); }
}

function refreshData(){
  ajax('api.php?a=list_services', function(r){
    var h=''; for(var i=0;i<r.length;i++){ var s=r[i];
      h+= '<div class="card"><b>#'+s.id+'</b> '+s.name+' ('+s.type+')<br>Créneaux: '+(s.slots||[]).join(', ')+'</div>';
    }
    document.getElementById('svc').innerHTML=h;
  });
  var em=(document.querySelector('input[name=mail]')||{}).value || (document.cookie.split('=')[1]||'user@example.com');
  ajax('api.php?a=list_bookings&e='+encodeURIComponent(em), function(b){
    var t=''; for(var j=0;j<b.length;j++){ var k=b[j];
      t+='<div>#'+k.id+' svc='+k.service+' @ '+k.slot+'</div>';
    }
    document.getElementById('bk').innerHTML=t || '(aucune)';
  });
}

function RefreshDATA(){ refreshData(); }

window.addEventListener('load', function(){
    window.refreshData(); 
});