<?php
    $input_password = "24";
  $admin_password_hash = "0e745304954233478109486485467426";
  
  echo "MD5 Hash:" . md5($input_password) . "\n";

  if (md5($input_password) == $admin_password_hash) {

    echo "[==] True\n";

  } else {
  
    echo "[==] False\n";
  
  }

  if (md5($input_password) === $admin_password_hash) {

    echo "[===] True\n";

  } else {
  
    echo "[===] False\n";
  
  }
?>