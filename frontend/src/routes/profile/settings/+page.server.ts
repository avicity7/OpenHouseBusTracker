import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { error, redirect } from "@sveltejs/kit";

export const load = async ({ fetch, locals }) => {
  try {
    const email = locals?.session?.Email;

    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/get-user-settings/${email}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch: ${response.statusText}`);
    }

    const UserSettings = await response.json();

    return {
      UserSettings
    };
  } catch (error) {
    console.error("Error fetching user settings:", error);
    return {
      status: 500,
      body: { error: 'Internal Server Error' }
    };
  }
};

export const actions = {
  updateUserSettings: async({ request}): Promise<void> =>{
    const form =await request.formData()

    const Name = form.get('name');
    let Contact = form.get('contact');
    const Email = form.get('email')

    if (typeof Contact === 'string') {
      Contact = Contact.replace(/\s/g, ''); // to remove whitespace in contact input
    }
    
    console.log("what is submitted", Name, Contact, Email)
    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/users/update-settings`, {
      method: "PUT",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
              Name, 
              Contact,
              Email
          })
    });
      
    console.log(`update user settings successful`)
    if (!response.ok) {
        console.log(error)
      throw new Error("Failed to update user settings");
    }

    redirect(301, '/profile')
  }
}

